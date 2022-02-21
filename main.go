package main

import (
    
    "fmt"
    _ "github.com/go-sql-driver/mysql"

    "todo/app/controllers"
    "todo/app/models"
    
)

func main() {
    
    fmt.Println(models.Db)

    controllers.StartMainServer()

   
	
}
