package main

import (
	"fmt"
	//_ "github.com/lib/pq"
	"net/http"
)

//type Test struct {
//	Result string `json:"result"`
//}
//
//var db *sql.DB
//var err error
//
//type Word struct {
//	English string `json:"english"`
//}
//
//func main() {
//
//	//connect to db
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		"localhost",
//		5432,
//		"postgres",
//		"root",
//		"dict_database")
//
//	db, err = sql.Open("postgres", psqlInfo)
//	if err != nil {
//		log.Error("error connecting to DB: %v\n", err)
//	} else {
//		log.Print("DB Connected successful")
//	}
//
//	db = db
//
//	// Creating context for http response/request using iris
//	app := iris.New()
//	app.Post("/", addWord)
//
//	err = app.Run(iris.Addr("localhost:8080"))
//
//	if err != nil {
//		log.Error("error starting server: %v\n", err)
//	}
//}
//
//func addWord(ctx iris.Context) {
//	var err error
//	var word Word
//	err = ctx.ReadJSON(&word)
//	sqlStatement := `INSERT INTO words (english) VALUES ($1)`
//	result, err := db.Exec(sqlStatement, word.English)
//	if err != nil {
//		log.Print(err)
//	}
//	_, err = ctx.JSON(Test{word.English})
//	result = result
//}

func main(){
	fmt.Println("Go Docker Tutorial")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println(w, "Hello world")
	})
}
