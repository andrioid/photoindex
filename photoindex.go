package photoindex

import (
	"time"
)

type photoindexer interface {
	save(p *Photo) (err error)
	get(id string) (p *Photo, err error)
	New() (p *Photo)
	Search(s string, offset, perpage int) (i Iterator, err error)
}

type Tag string
type User string

type Photo struct {
	id          string
	title       string
	description string
	path        string
	tags        []Tag
	created     time.Time
	owner       User

	idx photoindexer
}

// Call flow:
// - indexer implementation on call Search
// - search inits, connects to db, queries
// - stores rows pointer in a context struct
// - Search registers a callback function into the context struct
//
// - when next is called, the callback is called, calling next on rows
type Iterator interface {
	Next() (has_next bool)
	Value() (p *Photo)
}

func (p *Photo) save() error {
	return p.idx.save(p)
}

func (p *Photo) New() (photo *Photo) {
	return nil
}

func (p *Photo) Get(id string) (photo *Photo, err error) {
	return p.idx.get(id)
}
