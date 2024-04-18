package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/geoffrey-anto/course-downloader/parser"
	"github.com/google/uuid"
)

type Downloader struct {
	CourseName string
}

func getTempFileName() string {
	return strings.Split(uuid.New().String(), "-")[0]
}

func (d Downloader) DownloadVideo(video parser.Video, finished chan bool) {
	fmt.Printf("video.Name Started: %v\n", video.Name)
	client := http.Client{}

	fileNameAudio := getTempFileName() + ".mp4"
	fileNameVideo := getTempFileName() + ".mp4"

	// Download audio
	audioRequest, err := http.NewRequest("GET", video.AudioLink, nil)
	if err != nil {
		panic(err)
	}

	audioResponse, err := client.Do(audioRequest)
	if err != nil {
		panic(err)
	}

	defer audioResponse.Body.Close()

	// Download video
	videoRequest, err := http.NewRequest("GET", video.VideoLink, nil)
	if err != nil {
		panic(err)
	}

	videoResponse, err := client.Do(videoRequest)

	if err != nil {

		panic(err)
	}

	defer videoResponse.Body.Close()

	// Save audio
	audioFile, err := os.Create("tmp/" + fileNameAudio)
	if err != nil {
		panic(err)
	}

	defer audioFile.Close()

	_, err = io.Copy(audioFile, audioResponse.Body)

	if err != nil {
		panic(err)
	}

	// Save video
	videoFile, err := os.Create("tmp/" + fileNameVideo)
	if err != nil {
		panic(err)
	}

	defer videoFile.Close()

	_, err = io.Copy(videoFile, videoResponse.Body)

	if err != nil {
		panic(err)
	}

	e := exec.Command("ffmpeg", "-i", "tmp/"+fileNameVideo, "-i", "tmp/"+fileNameAudio, "-c:v", "copy", "-c:a", "aac", d.CourseName+"/"+video.Name+".mp4")

	e.Run()

	os.Remove("tmp/" + fileNameAudio)
	os.Remove("tmp/" + fileNameVideo)

	finished <- true
}
