package main

import (
	"flag"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	directory   = flag.String("directory", "./videos/", "Name of video file to upload")
	description = flag.String("description", "", "Video description")
	category    = flag.String("category", "17", "Video category")
	keywords    = flag.String("keywords", "", "Comma separated list of video keywords")
	privacy     = flag.String("privacy", "public", "Video privacy status")
)

func main() {

	flag.Parse()

	if *description == "" {
		log.Fatalf("You must provide a description for your videos!\n\nexample: ./youTubeUpload --description=\"New Jersey Soccer Event\"\n\n")
	}

	ts, err := getTokenSource(youtube.YoutubeUploadScope)

	if err != nil {
		log.Fatalf("error getting token source: %v", err)
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithTokenSource(ts), option.WithScopes(youtube.YoutubeUploadScope))

	if err != nil {
		log.Fatalf("error creating youTube service: %v", err)
	}

	uploadVideos(
		parseVideos(
			getVideos(*directory),
			*description,
			*category,
			*privacy,
			*keywords,
		),
		service,
	)
}
