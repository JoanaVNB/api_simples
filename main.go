package main

import(
	"github.com/joanavnb/api/routes"
	//"github.com/joanavnb/api/models"
	"github.com/joanavnb/api/database"
	//"github.com/gin-gonic/gin"
)


func main(){
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
