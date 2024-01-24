package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type MariaDB interface {
	Open() error
}

var DB *gorm.DB

func OpenDB() {
	var err error
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 사용자명:비밀번호@tcp(호스트:포트)/데이터베이스명?옵션
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	// 연결 테스트 - ping 요청 보내기
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("데이터베이스 연결 성공")
}

// func SetORM() error {

// 	return DB.AutoMigrate(&dal.Board{})
// }

// if err := config.OpenDB(); err != nil {
// 	log.Fatal("데이터베이스 연결 실패:", err)
// } else {
// 	log.Println("데이터베이스 연결 성공")
// }

// if err := config.SetORM(); err != nil {
// 	log.Fatal("모델 마이그레이션 실패:", err)
// } else {
// 	log.Println("모델 마이그레이션 성공")
// }
