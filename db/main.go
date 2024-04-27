package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버 로드
	"log"
)

func main() {
	// DB 연결 정보
	dbUser := "user"
	dbPass := "password"
	dbName := "dbname"
	dbHost := "mysql"
	dbPort := "3306"

	// DB 연결
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 테이블 조회 쿼리
	sql := "SELECT id, name, http_client FROM site" // 테이블 이름 변경 필요

	// 쿼리 실행
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 결과 처리
	for rows.Next() {
		var id int
		var name string
		var url string

		// 각 레코드 데이터 추출
		err := rows.Scan(&id, &name, &url)
		if err != nil {
			log.Fatal(err)
		}

		// 데이터 출력
		fmt.Printf("ID: %d, 이름: %s, 이메일: %s\n", id, name, url)
	}

	// 에러 처리
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
