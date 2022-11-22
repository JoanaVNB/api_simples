package models

import(
	"gorm.io/gorm"
	//"github.com/joanavnb/api/models"
)

type Aluno struct{
	gorm.Model
	Nome string `json:"nome"`
	CPF		string `json:"cpf"`
	RG	 string `json:"rg"`
}

