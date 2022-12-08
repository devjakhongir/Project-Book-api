package storage


import (
	"database/sql"


	"app/moduls"
)


func Create (db *sql.DB,  book moduls.Book) (string, error){
	var (
		id  string
		query string
	)

	query = `
	   insert into 
	    	books(name, price, author_name, publisher_date, page, janr)
	   values ($1, $2, $3, $4, $5, $6)
	   returning id
	`
	err := db.QueryRow(
		query,
		book.Name,
		book.Price,
		book.AuthorName,
		book.PublisherDate,
		book.Page,
		book.Janr,
	).Scan(&id)

	if err != nil{
		return "", err
	}

	return id, nil
}

func GetById(db *sql.DB, id string) (moduls.Book, error){
	var (
		book moduls.Book
		query string
	)


	query = `
		select 
			id,
			name,
			price,
            author_name,
            publisher_date,
			page,
			janr
		from
			books
		where id = $1
	`

	err := db.QueryRow(
		query,
		id,
	).Scan(
		&book.Id,
		&book.Name,
		&book.Price,
		&book.AuthorName,
		&book.PublisherDate,
		&book.Page,
		&book.Janr,
		
		
	)

	if err != nil {
		return moduls.Book{}, err
	}
	return book, nil
}

func GetList(db *sql.DB) ([]moduls.Book, error){
	var(
		books []moduls.Book
		query string
	)

	query = `
		select
		   id
		   name,
		   price,
		   author_name ,
		   publisher_date,
		   page,
		   janr 
		from
			books
	`
	rows,  err := db.Query(query)

	if err != nil{
		return []moduls.Book{}, err
	}

	for rows.Next() {
		var book moduls.Book

		err = rows.Scan(
			//&book.Id,
		    &book.Name,
		    &book.Price,
		    &book.AuthorName,
			&book.PublisherDate,
		    &book.Page,
		    &book.Janr,
		
		)

		if err != nil {
			return []moduls.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil

}