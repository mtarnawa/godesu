package main

import (
	"fmt"
	"github.com/lordrusk/godesu"
)

func main() {
	// Initialize the client
	gochan := godesu.New()

	// All Get... functions return err and a struct
	// Fetch all boards
	_, boards := gochan.GetBoards()

	for _, b := range boards.All {
		println(b.Board)
	}

	// Scope work to one board
	w := gochan.Board("w")
	// Return full catalog of a board
	_, catalog := w.GetCatalog()

	for _, page := range catalog.Pages {
		for _, thread := range page.Threads {
			println(thread.Name)
		}
	}

	// Return one page from standard view
	_, page := w.GetPage(1)
	for _, page := range page.All {
		for _, post := range page.Posts {
			println(post.Md5)
		}
	}

	// Get the selected thread
	_, thread := w.GetThread(1565459)
	images := thread.Images()
	for _, post := range thread.Posts {
		fmt.Println(images[post.No])
	}

	// Get images from a thread
	//images := thread.Images()
	//for _, image := range images {
	//		fmt.Printf(
	//			"url: %s\n"+
	//				"filename: %s\n"+
	//				"extension: %s\n"+
	//				"original filename: %s\n",
	//			image.URL, image.Filename, image.Extension, image.OriginalFilename)
	//	}
}
