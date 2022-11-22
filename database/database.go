package database

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"github.com/joanavnb/api/models"
)
var (
	DB  *gorm.DB
	err error
)
//O método db.AutoMigrate possui a funcionalidade de migrar modelos(structs) que estão em código Go 
//para o banco de dados. Os modelos são criados no banco de dados utilizando como base os dados 
//acessados a partir da referência de memória de uma ou mais structs que são passadas como parâmetro
// para o método db.AutoMigrate.
func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
			log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}