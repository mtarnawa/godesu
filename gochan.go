package godesu

import (
	"fmt"
)

type Gochan struct {
	c *client
}

type board struct {
	name string
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
		c: newClient(),
		}
	}

func (g *Gochan) Board(name string) *board {
	return &board{name, g}
}

func (b *board) GetThread(number int) (err error, thread Thread) {
	err = b.gochan.c.Get(
		fmt.Sprintf("/%s/thread/%d.json", b.name, number), &thread)
	return
}

func (b *board) GetPage(number int) (err error, model PaginatedThreads) {
	err = b.gochan.c.Get(
		fmt.Sprintf("/%s/%d.json", b.name, number), &model)
	model.board = b.name
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
	err, thread = pt.c.Board(pt.board).GetThread(number)
	thread.board = pt.board
	return
}

func (b *board) GetCatalog() (err error, model CatalogModel) {

	err = b.gochan.c.Get(fmt.Sprintf("/%s/catalog.json", b.name), &model.Pages)
	model.board = b.name
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
	err, thread = cm.c.Board(cm.board).GetThread(number)
	thread.board = cm.board
	return
}

func (g *Gochan) GetBoards() (err error, model Boards) {
	err = g.c.Get("/boards.json", &model)
	return
}

func (t *Thread) Images() (result Images) {
	for _, p := range t.Posts {
		if p.Tim > 0 {
			filename := fmt.Sprintf("%d%s", p.Tim, p.Ext)
			result = append(result,
				Image{
					URL:              IMG(fmt.Sprintf("/%s/%s", t.board, filename)),
					OriginalFilename: p.Filename + p.Ext,
					Filename:         filename,
					Extension:        p.Ext,
				})
		}
	}
	return
}