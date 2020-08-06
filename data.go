package data

import(
	"fmt"
	"time"
//	"math/rand"
	"log"
)

type Project struct{
	Title string
	ID int
	Date time.Time
	Note string
}
type projects [ ]Project

func (p projects) Search( id int)(string , time.Time,string){
//	p := projects{}
	for i:=0 ;i<len(p);i++{
		if p[i].ID ==id{
			return p[i].Title, p[i].Date,p[i].Note
		//	fmt.Println()
		}
	}
	return " ", time.Now()," "
}
// Add i think it need to be corrected	


func (p projects) Modify(id int, perm bool, title string, note string){
//	p:= projects{}
	switch perm {
	case true:
		for i := 0; i < len(p); i++ {
			if p[i].ID == id{
				p[i].Title = title
				p[i].Note = note
				p[i].Date = time.Now()
			}
		}
	case false:
		log.Fatal("no permitted")
	}
}
func (p projects)show(){
//	pi:= Project{}
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])	
	}
}
