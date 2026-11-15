package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Book struct{
	ID string
	Title string
	Author string
	Price float64 
}


var books= []Book{
{ID:"1",Title:"book1",Author:"author1",Price:65.2},
{ID:"2",Title:"book2",Author:"author2",Price:78.2},
{ID:"3",Title:"book3",Author:"author3",Price:95.2},
{ID:"4",Title:"book4",Author:"author4",Price:85.2},

}

func getBooks (c *gin.Context){

	c.IndentedJSON(http.StatusOK,gin.H{"data":books})
}

func postBook (c *gin.Context){

	var newBook Book

	if err:= c.BindJSON(&newBook); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"mesage":"Book is not valid","error":err.Error()})

		return
	}
	books=append(books,newBook)
c.JSON(http.StatusCreated,gin.H{"message":"create new book successfully","data":books})


}

func getBookByID(c *gin.Context){
id :=c.Param("id")

if id == ""{
	c.JSON(http.StatusBadRequest,gin.H{"message":"ID is not valid"})
return
}

for _,v := range books {
	if v.ID == id{
       c.IndentedJSON(http.StatusOK,gin.H{"data":v})
	   return
	}
}

c.IndentedJSON(http.StatusNotFound,gin.H{"message":"book not found"})
}

func editBook(c *gin.Context){

	var book Book
	id:=c.Param("id")

	if err := c.BindJSON(&book);err!=nil{
       c.JSON(http.StatusBadRequest,gin.H{"message":"request is not valid","err":err.Error()})
	   return
	}

	for i,b := range books{
		if b.ID != id{
			continue
		}
		
		if book.Title != "" {
 books[i].Title=book.Title
		}

			if book.Author != "" {
 books[i].Author=book.Author
		}
			if book.Price != 0 {
 books[i].Price=book.Price
		}
        
		 books[i].Price=book.Price
	     c.IndentedJSON(http.StatusOK,gin.H{"message":"Books updated successfully","data":books})
return
		

	}

	     c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not found","data":books})


}


func main(){
	router := gin.Default()

	router.GET("/books",getBooks)
	router.POST("/books",postBook)
	router.GET("/books/:id",getBookByID)
	router.PATCH("/books/:id",editBook)

	router.Run()
}