package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Film struct {
	name   string
	rating int64
}

type User struct {
	name  string
	films []Film
}

func main() {
	var users = make([]User, 1)
	count := 0
	file, error := os.Open("../Movie_Ratings.csv")
	if error != nil {
		panic(error)
	}
	r := csv.NewReader(bufio.NewReader(file))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		count2 := 0
		if count == 0 {
			count2++
			user := User{name: record[count2]}
			users = append(users, user)
		}

		if count > 0 {

		}

		count++
	}
}
