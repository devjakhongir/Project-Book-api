package main

import (
	"fmt"
	"log"

	"app/config"
	"app/pkg/db"
	"app/storage"
	//"app/moduls"
)

func main (){
	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil{
		panic(err)
	}

	defer db.Close()

	// u := moduls.Book{       
	//     Name           : "kitob",
	//     Price 			:  2300,
	//     AuthorName  	:"Abulla Qahharov",
	//     PublisherDate   : "2016",
	//     Page 			: 200,
	//     Janr			: "Qissa",
		
	// }


	// id, err := storage.Create(db, u)
	// if err != nil{
	// 	log.Fatalf("error whiling create user: %v", err)
	// }

	//  fmt.Println(id)

	// book, err := storage.GetById(db, id)
	// if err != nil{
	// 	log.Fatalf("error whiling create user: %v", err)
	// }
	// 	fmt.Println(book)
	books, err := storage.GetList(db)
	if err != nil{
		log.Fatalf("error whiling create user: %v", err)
	}
	fmt.Println(books)
 }
