// Package cmd is a package that contains subcommands for the reddit-downloader CLI command.
package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/sync/semaphore"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reddit-downloader ",
		Short: "reddit-downloader downloads images from Reddit.",
		Long: `reddit-downloader downloads images from Reddit.

You need to set the following environment variables:
  - GO_REDDIT_CLIENT_ID: to set the client's id.
  - GO_REDDIT_CLIENT_SECRET: to set the client's secret.
  - GO_REDDIT_CLIENT_USERNAME: to set the client's username.
  - GO_REDDIT_CLIENT_PASSWORD: to set the client's password.
`,
		RunE: download,
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmd.Flags().StringSliceP("sub-reddit", "s", []string{}, "Specify subreddit names to download images (Delimited by commas)")
	cmd.Flags().StringP("output", "o", "output", "Specify output directory to save images")

	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newBugReportCmd())
	return cmd
}

// download is a main function of reddit-downloader command.
func download(cmd *cobra.Command, _ []string) error {
	downloader, err := newDownloader(cmd)
	if err != nil {
		return err
	}
	return downloader.download()
}

// downloader is a main function of reddit-downloader command.
type downloader struct {
	*reddit.Client
	opt *option
}

// newDownloader returns a new downloader. It returns an error if the downloader is invalid.
func newDownloader(cmd *cobra.Command) (*downloader, error) {
	client, err := reddit.NewClient(reddit.Credentials{}, reddit.FromEnv)
	if err != nil {
		return nil, err
	}
	opt, err := newOption(cmd)
	if err != nil {
		return nil, err
	}
	return &downloader{
		Client: client,
		opt:    opt,
	}, nil
}

// download downloads images from Reddit.
func (d *downloader) download() error {
	postInSubReddits := make(map[string][]*reddit.Post)
	for _, v := range d.opt.subReddits {
		posts, err := d.getPosts(v)
		if err != nil {
			return err
		}
		postInSubReddits[v] = posts
	}

	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(int64(runtime.NumCPU()))
	for subRedditName, posts := range postInSubReddits {
		if err := os.MkdirAll(filepath.Join(d.opt.outputDir, subRedditName), 0777); err != nil {
			return err
		}

		for _, post := range posts {
			if post.URL == "" || post.IsSelfPost {
				continue
			}
			if !isURLImage(post.URL) {
				log.Info(fmt.Sprintf("skipped: Title:'%s' does not have image", post.Title))
				continue
			}

			wg.Add(1)
			go func(subRedditName string, post *reddit.Post) {
				defer wg.Done()

				if err := sem.Acquire(context.Background(), 1); err != nil {
					log.Error("failed to acquire semaphore token: %v\n", err)
					return
				}
				defer sem.Release(1)

				out := filepath.Join(d.opt.outputDir, subRedditName, generateFileName(post))
				err := downloadMedia(post.URL, out)
				if err != nil {
					log.Error("failed to download media: %v\n", err)
				}
				fmt.Printf("media downloaded: %s (%s)\n", post.Title, out)
			}(subRedditName, post)
		}
	}
	wg.Wait()
	return nil
}

func (d *downloader) getPosts(subRedditName string) ([]*reddit.Post, error) {
	posts, _, err := d.Subreddit.TopPosts(context.Background(), subRedditName, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 500,
		},
		Time: "month",
	})
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// option is a option for the reddit-downloader command.
type option struct {
	// subReddits is a subreddit name to download images.
	subReddits []string
	// outputDir is a directory to save images.
	outputDir string
}

// newOption returns a new option. It returns an error if the option is invalid.
func newOption(cmd *cobra.Command) (*option, error) {
	subReddit, err := cmd.Flags().GetStringSlice("sub-reddit")
	if err != nil {
		return nil, err
	}
	if len(subReddit) == 0 {
		return nil, errors.New("--sub-reddit option is required. please see help")
	}

	outputDir, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, err
	}

	return &option{
		subReddits: subReddit,
		outputDir:  outputDir,
	}, nil
}

// Execute run leadtime process.
func Execute() int {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		return 1
	}
	return 0
}

// generateFileName generates a file path for the image.
func generateFileName(post *reddit.Post) string {
	ext := filepath.Ext(post.URL)
	title := strings.Replace(post.Title, " ", "_", -1)
	title = strings.Replace(title, "/", "_", -1)
	filename := fmt.Sprintf("%s_%s%s", post.ID, title, ext)
	return filename
}

// downloadMedia downloads the media from the URL and saves it to the filepath.
func downloadMedia(url, filepath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

// isURLImage returns true if the URL is an image.
func isURLImage(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	return strings.Contains(contentType, "image")
}
