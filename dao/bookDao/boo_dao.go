package bookDao

import (
	"fmt"
	"github.com/uuidcode/echo-test2/dao"
	"github.com/uuidcode/echo-test2/domain"
	"github.com/uuidcode/echo-test2/util"
)

func FindAll() []domain.Book {
	db := dao.ConnectDB()

	sql := `
	select * from book
	
	`

	rows, err := db.Queryx(sql)
	util.CheckErr(err)

	defer rows.Close()
	defer db.Close()

	var bookList []domain.Book

	for rows.Next() {
		var currentBook domain.Book
		err := rows.StructScan(&currentBook)
		util.CheckErr(err)

		bookList = append(bookList, currentBook)
	}

	for i, currentBook := range bookList {
		fmt.Printf("%v %+v\n", i, currentBook)
	}

	return bookList
}
