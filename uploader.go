package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"google.golang.org/api/youtube/v3"
)

// youTubeVideo is a struct that defines the youTube video object
type youTubeVideo struct {
	File        string
	Title       string
	Description string
	CategoryId  string
	Privacy     string
	Keywords    string
}

// contains checks if an element exists in given slice
func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// getVideos gets the videos from a given directory and adds them to a channel
func getVideos(directory string) <-chan string {
	c := make(chan string)
	go func() {
		files, err := ioutil.ReadDir(directory)

		if err != nil {
			log.Fatal(err)
		}

		extensions := []string{".flv", ".mp4", ".mov", ".mpeg-1", ".mpeg-2", ".mpeg4", ".mpg", ".avi", ".wmv", ".mpegps", ".3gpp", ".webm", ".dnxhr", ".prores", ".cineform", ".hevc", ".mts"}
		for _, file := range files {
			ext := filepath.Ext(file.Name())
			if contains(extensions, strings.ToLower(ext)) {
				c <- directory + file.Name()
			}
		}

		close(c)
	}()

	return c
}

// parseVideos takes the videos and instantiates a new struct
func parseVideos(c <-chan string, description string, category string, privacy string, keywords string) <-chan *youTubeVideo {
	vc := make(chan *youTubeVideo)
	go func() {
		for line := range c {
			_, filename := filepath.Split(line)
			t := strings.TrimSuffix(filename, filepath.Ext(filename))
			re := strings.NewReplacer("_", " ", ":", "/")
			title := re.Replace(t)

			video := &youTubeVideo{
				File:        line,
				Title:       title,
				Description: description,
				CategoryId:  category,
				Privacy:     privacy,
				Keywords:    keywords,
			}

			vc <- video
		}

		close(vc)
	}()

	return vc
}

// createVideo instants a new youTube.Video object for uploading
func createVideo(video *youTubeVideo) *youtube.Video {
	yt := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       video.Title,
			Description: video.Description,
			CategoryId:  video.CategoryId,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: video.Privacy},
	}

	if strings.Trim(*keywords, "") != "" {
		yt.Snippet.Tags = strings.Split(video.Keywords, ",")
	}

	return yt

}

// uploadVideos uploads each video from the given channel
func uploadVideos(vc <-chan *youTubeVideo, service *youtube.Service) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for video := range vc {
			uploadVideo(video, service)
		}
	}()
	wg.Wait()
}

// uploadVideo uploads the individual video
func uploadVideo(video *youTubeVideo, service *youtube.Service) {
	upload := createVideo(video)

	call := service.Videos.Insert([]string{"snippet,status"}, upload)

	file, err := os.Open(video.File)

	if err != nil {
		log.Fatalf("Error opening %v: %v", video.File, err)
	}

	defer file.Close()

	response, err := call.Media(file).Do()
	handleError(err, "")
	fmt.Printf("Upload of %v successful! Video ID: %v\n", upload.Snippet.Title, response.Id)
}
