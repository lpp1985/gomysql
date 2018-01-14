package main

import (
	"fmt"
	"gomysql"
)

func main() {
	c, err := gomysql.SetConfig("database.ini","DB","Taxon")
	if err != nil {
		fmt.Println(err)
	}
	t:=c.SetTable("TaxonName")
        data := t.Fileds("taxon", "name").Where("taxon <1000").FindAll() 
	//fmt.Println(data)
	for _,data:=range data{
		fmt.Println( data["name"] )
}

}
