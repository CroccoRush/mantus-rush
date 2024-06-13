package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

// Есть две БД: PostgreSQL, Redis.
// PostgreSQL хранит в себе полную информацию о всех экспериментах и шаблонах, которые есть сейчас и были в прошлом.
// В том числе файлы конфигурация, входные данные и структуру конвейера.
// Redis хранит информацию о запущенных и недавно завершенных экспериментах: время запуска, входные данные, выходные данные, статус.

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBTimezone string
	DBName     string
	DB         *gorm.DB
	sqlDB      *sql.DB
	err        error
	Schemas    map[string]interface{}
)

func init() {
	DBHost = "localhost"
	DBPort = "5432"
	DBUser = "postgres"
	DBPassword = "159951"
	DBTimezone = "Europe/Moscow"
	DBName = "eris"

	Schemas = map[string]interface{}{
		"user":          User{},
		"module":        Module{},
		"template":      Template{},
		"experiment":    Experiment{},
		"resultMessage": ResultMessage{},
	}

}

func Run(globsig chan bool, wg *sync.WaitGroup, onStartWg *sync.WaitGroup) {

	defer wg.Done()
	innerch := make(chan bool)

	createDatabase()
	if sqlDB, err = DB.DB(); err != nil {
		log.Fatal(err)
	}
	createTables()
	createSuperUser()

	onStartWg.Done()
	utils.Dog(globsig, innerch, sqlDB, "database")
}

func createDatabase() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		DBHost, DBPort, DBUser, DBPassword, DBName, DBTimezone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s",
			DBHost, DBPort, DBUser, DBPassword, DBTimezone)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			log.Println(err)
		}

		DB.Exec(
			fmt.Sprintf("CREATE DATABASE %s", DBName),
		)
		sqlDB, err = DB.DB()
		if err != nil {
			log.Fatalln(err)
		}
		sqlDB.Close()

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			DBHost, DBPort, DBUser, DBPassword, DBName, DBTimezone)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func createTables() {

	for name, schema := range Schemas {
		if !DB.Migrator().HasTable(&schema) {
			if err := DB.Migrator().CreateTable(&schema); err != nil {
				log.Fatalln(err)
			}
			log.Printf("table -----%s----- successfully created", name)
		}
	}
}

func createSuperUser() {

	var superuser User
	DB.Find(&superuser, User{ID: 1})
	if superuser.Login != "" && superuser.Login != "SUPER" {
		log.Fatalln("wrong superuser")
	} else if superuser.Login == "" {
		user := User{
			ID:       1,
			Login:    "SUPER",
			Password: "SUPER",
			Email:    "SUPER",
			Role:     2,
		}

		result := DB.Create(&user)
		if result.Error != nil {
			log.Fatalf("add record error: %s", result.Error)
			return
		}

		log.Printf("the superuser was successfully added. ID: 1")
	} else {
		log.Printf("the superuser exist already. ID: 1")
	}
}
