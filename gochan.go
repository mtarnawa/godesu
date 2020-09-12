package godesu

import (
	"fmt"
	"strconv"
)

type Gochan struct {
	c *Client
}

type board struct {
	Name string
	gochan *Gochan
}

type Images []Image

type Image struct {
	URL string
	OriginalFilename string
	Filename string
	Extension string
}

func New() *Gochan {
	return &Gochan{
		c: NewClient(),
		}
	}

func (g *Gochan) Board(name string) *board {
	return &board{name, g}
}

func (b *board) GetThread(number int) (err error, thread Thread) {
	err = b.gochan.c.Get(
		fmt.Sprintf("/%s/thread/%d.json", b.Name, number), &thread)
	thread.Board = b.Name
	thread.Posts = make(map[int]Post)
	for _, post := range thread.posts {
		thread.Posts[post.No] = post
	}
	return
}

func (b *board) GetPage(number int) (err error, model PaginatedThreads) {
	err = b.gochan.c.Get(
		fmt.Sprintf("/%s/%d.json", b.Name, number), &model)
	model.Board = b.Name
	model.c = b.gochan
	return
}

func (pt *PaginatedThreads) Thread(number int) (err error, thread Thread) {
	var found bool
	pages:
	for _, page := range pt.All {
		for _, post := range page.Posts {
			if post.No == number {
				found = true
				break pages
			}
		}
	}
	if !found {
		return ErrThreadNotFound{number}, Thread{}
	}
	err, thread = pt.c.Board(pt.Board).GetThread(number)
	thread.Board = pt.Board
	return
}

func (b *board) GetCatalog() (err error, model CatalogModel) {

	err = b.gochan.c.Get(fmt.Sprintf("/%s/catalog.json", b.Name), &model.Pages)
	model.Board = b.Name
	model.c = b.gochan
	return
}

func (cm *CatalogModel) Thread(number int) (err error, thread Thread) {
	var found bool
	pages:
	for _, p := range cm.Pages {
		for _, thread := range p.Threads {
			if thread.No == number {
				found = true
				break pages
			}
		}
	}
	if !found {
		return ErrThreadNotFound{number}, Thread{}
	}
	err, thread = cm.c.Board(cm.Board).GetThread(number)
	thread.Board = cm.Board
	return
}

func (g *Gochan) GetBoards() (err error, model Boards) {
	err = g.c.Get("/boards.json", &model)
	return
}

func (t *Thread) Images() (result Images) {
	for _, p := range t.Posts {
		if p.Tim > 0 {
			result = append(result,
				Image{
					URL:              IMG(fmt.Sprintf("/%s/%s", t.Board, strconv.FormatInt(p.Tim, 10)+p.Ext)),
					OriginalFilename: p.Filename + p.Ext,
					Filename:         strconv.FormatInt(p.Tim, 10),
					Extension:        p.Ext,
				})
		}
	}
	return
}
