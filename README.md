# YouTubeUploader
This application allows for the user to upload videos to YouTube via the YouTube API <https://developers.google.com/youtube/v3/docs/videos/insert>

## Building
As a go application, you will need to build based upon the environment it will be used in. For reference, you may use `$ go build` when running in your own environment.

## Authentication
YouTube prefers using the oath2 standard. In order to authenticate properly, the following steps must be implemented.

1. change `example_client_secret.json` to `client_secret.json`
2. Enter your Google Cloud credentials
3. Run the application (see below)

## Installation
1. Create a folder to put the `youTubeUpload` executable file.
2. Create a folder inside the folder from step `1.` called `videos`
3. Inside the folder from step `1.` run `$ chmod u+x youTubeUpload`
4. Run the application by following the steps in **Running** below.  
--> If running on a Mac, You may be asked to give permissions to the application. Directions can be found at <https://support.apple.com/en-us/HT202491>. _Note that you may need to run the application again after updating permissions._  

## Running
Follow the below steps to upload the videos. Note that there is a constraint based upon your account limits.
1. Move all videos to be uploaded into the `videos` folder.
2. Run `$ ./youTubeUpload` if running the binary or if developing `$ go run run.go errors.go oauth2.go uploader.go`  
_You may need to authenticate. Follow directions in the terminal and on the webpage presented by Google._  

The following arguments are allowed.

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
