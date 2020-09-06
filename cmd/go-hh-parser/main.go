package main

import (
	"go-hh-parser/internal/app/core"
	"log"
)

func main() {
	core := core.New()

	urls := []string{
		"https://api.hh.ru/vacancies/38468170",
		"https://api.hh.ru/vacancies/38468172",
		"https://api.hh.ru/vacancies/38468173",
	}

	for _, url := range urls {

		vacancy, err := core.Req.GetVacancy(url)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(vacancy.Data)
	}
}
