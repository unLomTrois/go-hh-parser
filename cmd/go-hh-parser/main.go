package main

import (
	"go-hh-parser/internal/app/core"
	"log"
)

func main() {
	core := core.New()

	vacancy, err := core.Req.GetVacancy("https://api.hh.ru/vacancies/38468170")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(vacancy.Data)

}
