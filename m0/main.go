package main

import (
	"fmt"
	"github.com/mirrordust/splendour/m0/repo"
)

func main() {
	//web.Server().Run()
	var posts []repo.Post
	err := repo.FindAll(&posts, repo.Condition{
		Query:  nil,
		Args:   nil,
		Orders: nil,
		Offset: 0,
		Limit:  0,
	})
	fmt.Println("===")
	fmt.Println(posts)
	fmt.Println(err)
}

// func testDB() {
// 	fmt.Println("====== create ======")
// 	tag := Tag{Name: "tag1", Code: 0b1}
// 	tags := []Tag{
// 		{Name: "tag2", Code: 0b10},
// 		{Name: "tag3", Code: 0b100},
// 		{Name: "tag4", Code: 0b1000},
// 		{Name: "tag5", Code: 0b10000},
// 		{Name: "tag6", Code: 0b100000},
// 	}
// 	Create(&tag)
// 	Create(&tags)
// 	fmt.Printf("tag: %+v\n", tag)
// 	fmt.Println("tags: ")
// 	for _, tag_ := range tags {
// 		fmt.Printf("%+v\n", tag_)
// 	}

// 	fmt.Println("====== retrieve ======")
// 	var tag2 Tag
// 	var tag3 Tag
// 	var tags2 []Tag
// 	var tags3 []Tag
// 	FindOne(&tag2, 2)
// 	FindOne(&tag3, 3)
// 	FindAll(&tags2, "1=1")
// 	FindAll(&tags3, "id in ?", []uint64{1, 5})
// 	fmt.Printf("tag2: %+v\n", tag2)
// 	fmt.Printf("tag3: %+v\n", tag3)
// 	fmt.Println("tags2: ")
// 	for _, tag_ := range tags2 {
// 		fmt.Printf("%+v\n", tag_)
// 	}
// 	fmt.Println("tags3: ")
// 	for _, tag_ := range tags3 {
// 		fmt.Printf("%+v\n", tag_)
// 	}

// 	fmt.Println("====== update ======")
// 	tag2.Name = "tag999改"
// 	fmt.Printf("tag2: %+v\n", tag2)
// 	UpdateOne(&tag2, tag2)

// 	fmt.Println("====== delete ======")
// 	Delete(&tag3)
// 	Delete(&tags3)
// }
