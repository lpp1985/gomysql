package gomysql

import (
	"log"
	"testing"
)

func TestMysql(*testing.T) {
	c, err := SetConfig("general.ini","DB","Taxon")
	if err != nil {
		log.Println(err)
	}
	//data := c.SetTable("user").FindOne()
	log.Println(c)
}
