package photoindex

import (
	"time"
)

type photoindexer interface {
	save(p *Photo) (err error)
	//Get() (p *Photo)
	New() (p *Photo)
}

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

type Tag string

type User string

func (p *Photo) save() error {
	return p.idx.save(p)
}

func (p *Photo) New() (photo *Photo) {
	return nil
}

func (p *Photo) Get(id string) (photo Photo, err error) {
	photo = Photo{}
	return photo, nil
}
