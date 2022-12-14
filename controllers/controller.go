package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joanavnb/api/models"
	"github.com/joanavnb/api/database"
)

func ExibeTodosAlunos(c *gin.Context){
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context){
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E aí" + nome + ", tudo beleza?",
	})
}

func CriaNovoAluno(c *gin.Context){
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorID(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName(("id"))

	database.DB.First(&aluno, id)
	if aluno.ID == 0{
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
	return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)
	if aluno.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",})
	} 
	if aluno.ID == 0  {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
	}
}	

func EditaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)
//se err não for nil{}
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context){
	var aluno models.Aluno
	cpf := c.Param("cpf")
				
//do modelo de aluno, busque o primeiro com o cpf requisitado
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0  {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}