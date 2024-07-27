// Description: This is the main file that will be used to run the image processing pipeline.
// It will read images from the images folder, resize them, convert them to grayscale,
// and save them to the output folder.
//
// NOTE: Referenced and copied from `https://github.com/code-heim/go_21_goroutines_pipeline`
//
// - This is a learning exercise to understand how to use goroutines and channels to create a pipeline.
//
// Additional Reference:
//
// - https://www.youtube.com/watch?v=8Rn8yOQH62k
package main

import (
	"fmt"
	"image"
	"log"
	"strings"

	"github.com/veegrace/playground/golang/concurrency/image/processor"
)

type job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan job {
	out := make(chan job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			job := job{
				InputPath: p,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
			}

			job.Image = processor.ReadImage(p)
			out <- job
		}
		log.Print("closing loadImage")
		close(out)
	}()
	return out
}

func resize(in <-chan job) <-chan job {
	out := make(chan job)
	go func() {
		for job := range in { // Read from the input channel
			job.Image = processor.Resize(job.InputPath, job.Image)
			out <- job
		}
		log.Print("closing resize")
		close(out)
	}()
	return out
}

func convertToGrayscale(in <-chan job) <-chan job {
	out := make(chan job)
	go func() {
		for job := range in { // Read from the input channel
			job.Image = processor.Grayscale(job.InputPath, job.Image)
			out <- job
		}
		log.Print("closing convertToGrayscale")
		close(out)
	}()
	return out
}

func saveImage(in <-chan job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range in { // Read from the input channel
			processor.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		log.Print("closing saveImage")
		close(out)
	}()
	return out
}

func main() {
	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
	// Significance of the closing out channel of each stage,
	// channel is closed individually right after each stage is done processing.
	//
	// ref: close(out)
	// each stage in the pipeline.
}
