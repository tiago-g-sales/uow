package main

import (
	"context"
	"database/sql"

	"github.com/tiago-g-sales/uow/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil{
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)


	listCategories(ctx, queries)

}
func createCategory(ctx context.Context, queries *db.Queries){

	err := queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: "Backend",
	} )

	if err != nil {
		panic(err)
	}
}

func listCategories(ctx context.Context, queries *db.Queries){

	if true {
		panic("")
	}	

}



