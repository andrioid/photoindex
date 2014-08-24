package photoindex

import (
	"database/sql"
	"fmt"

	//"log"
	uuid "code.google.com/p/go-uuid/uuid"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

type Sqlite struct {
	db       *sql.DB
	filename string
}

func (d *Sqlite) New() (p *Photo) {
	photo := &Photo{}
	photo.idx = d
	photo.created = time.Now()
	return photo
}

func (d *Sqlite) save(p *Photo) (err error) {
	if err = d.init(); err != nil {
		return err
	}

	if p.id == "" {
		// Insert
		err = d.insert(p)
	} else {
		err = d.update(p)
	}
	return err
}

func (d *Sqlite) insert(p *Photo) (err error) {
	var db = d.db
	if d.db == nil {
		return fmt.Errorf("Database not initialized")
	}
	var newId = uuid.New()

	fmt.Println(db)
	_, err = db.Exec("INSERT INTO photos (id, title, description, path, owner, created) VALUES (?, ?, ?, ?, ?, ?)", newId, p.title, p.description, p.path, string(p.owner), p.created)
	if err != nil {
		return err
	}
	p.id = newId
	return nil
}

func (d *Sqlite) update(p *Photo) (err error) {
	var db = d.db
	// mooo
	_, err = db.Exec("UPDATE photos SET title=?, description=?, path=?, owner=?, created=?", p.title, p.description, p.path, string(p.owner), p.created)
	return err
}

func (d *Sqlite) get() (p *Photo, err error) {
	photo := &Photo{}
	return photo, err
}

func (d *Sqlite) init() (err error) {
	if d.filename == "" {
		return fmt.Errorf("No filename specified for context.")
	}
	if d.db == nil {
		d.db, err = sql.Open("sqlite3", d.filename)
		if err != nil {
			return err
		}

		if _, err := os.Stat(d.filename); os.IsNotExist(err) {
			fmt.Printf("Creating DB...\n")
			if _, err = d.db.Exec("CREATE TABLE photos (id text not null primary key, title text, description text, path text, owner text, created datetime)"); err != nil {
				return err
			}

			if _, err = d.db.Exec("CREATE TABLE photo_tags (tag string not null, photo_id string not null)"); err != nil {
				return err
			}

			if _, err = d.db.Exec("CREATE INDEX combo ON photo_tags (tag, photo_id)"); err != nil {
				return err
			}
		}
	}
	return nil
}

/*
func (d *Sqlite) Search(conditions string) (photos []Photo, err error) {

}
*/
