package photoindex

import (
	"fmt"
	"testing"
)

func TestMoo(t *testing.T) {
	db := Sqlite{filename: "index.db"}
	photo := db.New()
	photo.title = "oh, its pie"
	photo.description = "nothing"
	photo.owner = "nobody"
	photo.path = "data/images/awesomepic.jpg"
	err := photo.save()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}
