# reddit-downloader - Reddit images downloader
`reddit-downloader` is a CLI command that downloads images from posts within a subreddit. It allows you to specify multiple subreddits in a single execution. Thanks to parallelization, downloading with `reddit-downloader` enables you to gather images quickly.


## How to install
### Use "go install"
If you does not have the golang development environment installed on your system, please install golang from the [golang official website](https://go.dev/doc/install)

```
$ go install github.com/nao1215/reddit-downloader@latest
```

## How to use
### Environment variables for using reddit-downloader
reddit-downloader uses the following environment variables to access reddit API. If you get client id and client secret, you can create a client from [reddit apps](https://www.reddit.com/prefs/apps).
- GO_REDDIT_CLIENT_ID : client's id.
- GO_REDDIT_CLIENT_SECRET : client's secret.
- GO_REDDIT_CLIENT_USERNAME : client's username.
- GO_REDDIT_CLIENT_PASSWORD : client's password.

### Download images from subreddit
The `reddit-downloader` command accepts multiple subreddits through the `--sub-reddit` option and saves the images in the directory `output/${sub-reddit-name}`. The destination directory `output` can be changed using the `--output` option.
```
$ export GO_REDDIT_CLIENT_ID=xxxxxxxxxxxxxx ※ set your client id
$ export GO_REDDIT_CLIENT_SECRET=xxxxxxxxxxxxxx ※ set your client secret
$ export GO_REDDIT_CLIENT_USERNAME=xxxxxxxxxxxxxx ※ set your client username
$ export GO_REDDIT_CLIENT_PASSWORD=xxxxxxxxxxxxxx ※ set your client password

$ reddit-downloader --sub-reddit=wallpaper,MobileWallpaper
fetching posts from 'wallpaper' sub reddit
fetching posts from 'MobileWallpaper' sub reddit

media downloaded: Chinese poster Spider-Man: Across the Spider-Verse (output/MobileWallpaper/13sh0r9_Chinese_poster_Spider-Man:_Across_the_Spider-Verse.jpg)
skipped: Title:'Some from my collection of wallpapers': not support reddit gallery (URL:https://www.reddit.com/gallery/13py2gm)
media downloaded: Memories (output/MobileWallpaper/13hggwb_Memories.jpg)
media downloaded: Will be rocking for months to come (output/MobileWallpaper/13xbchy_Will_be_rocking_for_months_to_come.jpg)
media downloaded: Meteor Strike. (output/MobileWallpaper/13waxm0_Meteor_Strike..jpg)
 :
 :
```

If a post does not contain an image, it will be skipped. Furthermore, it is not possible to download the images from the gallery.
```
2023/06/11 16:05:07 INFO skipped: Title:'Bad post' does not have image
```

## Contributing
Contributions are welcome. Contributions are not only related to development. For example, GitHub Star motivates me to develop!

## Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/reddit-downloader/issues)
You can use the bug-report subcommand to send a bug report.

```
$ reddit-downloader bug-report
※ Open GitHub issue page by your default browser
```

## LICENSE
The reddit-downloader project is licensed under the terms of [MIT LICENSE](./LICENSE)