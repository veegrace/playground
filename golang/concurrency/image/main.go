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

func loadImage(done <-chan interface{}, paths []string) <-chan job {
	out := make(chan job)
	go func() {
		// This is to ensure that the channel is closed after all the jobs are done.
		defer func() {
			log.Print("closing loadImage")
			close(out)
		}()

		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			select {
			case <-done:
				return
			default:
				job := job{
					InputPath: p,
					OutPath:   strings.Replace(p, "images/", "images/output/", 1),
				}

				job.Image = processor.ReadImage(p)
				out <- job
			}
		}
	}()
	return out
}

func resize(done <-chan interface{}, in <-chan job) <-chan job {
	out := make(chan job)
	go func() {
		// This is to ensure that the channel is closed after all the jobs are done.
		defer func() {
			log.Print("closing resize")
			close(out)
		}()

		for job := range in { // Read from the input channel
			select {
			case <-done:
				return
			default:
				job.Image = processor.Resize(job.InputPath, job.Image)
				out <- job
			}
		}
	}()
	return out
}

func convertToGrayscale(done <-chan interface{}, in <-chan job) <-chan job {
	out := make(chan job)
	go func() {
		// This is to ensure that the channel is closed after all the jobs are done.
		defer func() {
			log.Print("closing convertToGrayscale")
			close(out)
		}()

		for job := range in { // Read from the input channel
			select {
			case <-done:
				return
			default:
				job.Image = processor.Grayscale(job.InputPath, job.Image)
				out <- job
			}
		}
	}()
	return out
}

func saveImage(done <-chan interface{}, in <-chan job) <-chan bool {
	out := make(chan bool)
	go func() {
		// This is to ensure that the channel is closed after all the jobs are done.
		defer func() {
			log.Print("closing saveImage")
			close(out)
		}()

		for job := range in { // Read from the input channel
			select {
			case <-done:
				return
			default:
				processor.WriteImage(job.OutPath, job.Image)
				out <- true
			}
		}
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

	done := make(chan interface{})
	defer close(done)

	channel1 := loadImage(done, imagePaths)
	channel2 := resize(done, channel1)
	channel3 := convertToGrayscale(done, channel2)
	writeResults := saveImage(done, channel3)

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
