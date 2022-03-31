# YouTubeUploader
This application allows for the user to upload videos to YouTube via the YouTube API <https://developers.google.com/youtube/v3/docs/videos/insert>

## Building
As a go application, you will need to build based upon the environment it will be used in. For reference, you may use `$ go build` when running in your own environment.

## Authentication
YouTube prefers using the oath2 standard. In order to authenticate properly, the following steps must be implemented.

1. change `example_client_secret.json` to 'client_secret.json'
2. Enter your Google Cloud credentials
3. Run the applicatin (seel below)

## Running
After account authentication (oauth2), follow the below steps to upload the videos. Note that there is a constraint based upon your account limits.
1. Move all videos into the video 
2. Run `$ ./youTubeUpload`. The following arguments are allowed
```
-category string
      Video category (default "17")
-description string
      Video description
-directory string
      Name of video file to upload (default "./videos/")
-keywords string
      Comma separated list of video keywords
-privacy string
      Video privacy status (default "public")
```
