# reddit-dl - Reddit images downloader
`reddit-dl` is a CLI command that downloads images from posts within a subreddit. It allows you to specify multiple subreddits in a single execution. Thanks to parallelization, downloading with `reddit-dl` enables you to gather images quickly.


## How to install
### Use "go install"
If you does not have the golang development environment installed on your system, please install golang from the [golang official website](https://go.dev/doc/install)

```
$ go install github.com/nao1215/reddit-dl@latest
```

## How to use
### Environment variables for using reddit-dl
reddit-dl uses the following environment variables to access reddit API. If you get client id and client secret, you can create a client from [reddit apps](https://www.reddit.com/prefs/apps).
- GO_REDDIT_CLIENT_ID : client's id.
- GO_REDDIT_CLIENT_SECRET : client's secret.
- GO_REDDIT_CLIENT_USERNAME : client's username.
- GO_REDDIT_CLIENT_PASSWORD : client's password.

### Download images from subreddit
The `reddit-dl` command accepts multiple subreddits through the `--sub-reddit` option and saves the images in the directory `output/${sub-reddit-name}`. The destination directory `output` can be changed using the `--output` option.
```
$ export GO_REDDIT_CLIENT_ID=xxxxxxxxxxxxxx ※ set your client id
$ export GO_REDDIT_CLIENT_SECRET=xxxxxxxxxxxxxx ※ set your client secret
$ export GO_REDDIT_CLIENT_USERNAME=xxxxxxxxxxxxxx ※ set your client username
$ export GO_REDDIT_CLIENT_PASSWORD=xxxxxxxxxxxxxx ※ set your client password

$ reddit-dl --sub-reddit=wallpaper,MobileWallpaper
media downloaded: Katana and a pistol [3840x2160] (output/wallpaper/13vq8r7_Katana_and_a_pistol_[3840x2160].jpg)
media downloaded: Rayquaza - Dragon Ascent [3440x1440] (output/wallpaper/13u8tbe_Rayquaza_-_Dragon_Ascent_[3440x1440].jpg)
media downloaded: May cause addiction [1600x1200] (output/wallpaper/1419ygj_May_cause_addiction_[1600x1200].png)
media downloaded: Amazing Islands in Ocean [1920 x 1200] (output/wallpaper/13y2kim_Amazing_Islands_in_Ocean_[1920_x_1200].jpg)
media downloaded: Miles Morales and Spider-Gwen [1920x1080] (output/wallpaper/1420iyi_Miles_Morales_and_Spider-Gwen_[1920x1080].jpg)
media downloaded: Anime girl in bar [3840x2160] (output/wallpaper/144u2eu_Anime_girl_in_bar_[3840x2160].jpg)
media downloaded: cat waiting for a train [1920x1080] (output/wallpaper/13jfg7t_cat_waiting_for_a_train_[1920x1080].jpg)
 :
 :
```

If a post does not contain an image, it will be skipped.
```
2023/06/11 16:05:07 INFO skipped: Title:'Bad post' does not have image
```

## Contributing
Contributions are welcome. Contributions are not only related to development. For example, GitHub Star motivates me to develop!

## Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/reddit-downloader/issues)
You can use the bug-report subcommand to send a bug report.

$ reddit-dl bug-report
※ Open GitHub issue page by your default browser

## LICENSE
The reddit-downloader project is licensed under the terms of [MIT LICENSE](./LICENSE)