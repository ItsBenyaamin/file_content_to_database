package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)


func main() {
	startTime := time.Now()

	checkFileExist()
	db := initDatabase()
	file := readFile()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			level, answers := sliceLine(scanner.Text())
			insetIntoDatabase(db, level, answers)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Printf("its took %s", elapsed)
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func checkFileExist() {
	if _, err := os.Stat("amirza.txt"); err != nil {
		log.Panic("file does not exist!")
	}
}

func initDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./amirza.db")
	checkErr(err)
	return db
}

func readFile() *os.File {
	myFile, err := os.Open("amirza.txt")
	checkErr(err)
	return myFile
}

func sliceLine(line string) (int64, []string) {
	lineSlice := strings.Split(line, "-")
	level := lineSlice[0]
	answers := strings.Split(lineSlice[1], "،")
	fmt.Println("Level is : ", level)
	fmt.Println("Answers is : ", strings.Join(answers, " "))

	return convertFaNumber(level), answers
}

func insetIntoDatabase(db *sql.DB, level int64, answers []string){
	for _,item := range answers{
		insertPrepare, err := db.Prepare("insert into am_w (level, word) values(?, ?)")
		checkErr(err)
		result ,err := insertPrepare.Exec(level, strings.TrimSpace(item))
		checkErr(err)
		fmt.Println("result : " , result)
	}
}

func convertFaNumber(level string) int64 {
	var wrappedLevel = ""
	for _, num := range strings.TrimSpace(level){
		switch string(num) {
		case "۱": wrappedLevel += "1"
		case "۲": wrappedLevel += "2"
		case "۳": wrappedLevel += "3"
		case "۴": wrappedLevel += "4"
		case "۵": wrappedLevel += "5"
		case "۶": wrappedLevel += "6"
		case "۷": wrappedLevel += "7"
		case "۸": wrappedLevel += "8"
		case "۹": wrappedLevel += "9"
		case "۰": wrappedLevel += "0"
		}
	}
	wrappedInt, err := strconv.ParseInt(wrappedLevel, 10, 64)
	checkErr(err)
	return wrappedInt
}