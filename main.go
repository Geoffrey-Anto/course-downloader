package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/geoffrey-anto/video-downloader/downloader"
	"github.com/geoffrey-anto/video-downloader/parser"
)

const FILE_NAME = "links.txt"
const COURSE_NAME = "Learn Ansible Basics - Beginners Course"

func main() {
	t := time.Now()

	parser := parser.Parser{}

	inputStream, err := os.Open(COURSE_NAME + "/" + FILE_NAME)
	os.Mkdir(COURSE_NAME, 0755)

	if err != nil {
		panic(err)
	}

	reader := io.Reader(inputStream)

	videos, err := parser.ParseFile(&reader)

	downloader := downloader.Downloader{CourseName: COURSE_NAME}

	finished := make(chan bool)

	for video := range videos {
		go downloader.DownloadVideo(videos[video], finished)
	}

	for i := 0; i < len(videos); i++ {
		a := <-finished

		if a {
			fmt.Printf("Downloaded %v\n", videos[i].Name)
		}
	}

	fmt.Printf("Time taken: %v\n", time.Since(t))

	if err != nil {
		panic(err)
	}
}
