package main

import (
	"go-hh-parser/internal/app/core"
	"log"
	"sync"
)

func main() {
	core := core.New()

	var wg sync.WaitGroup

	urls := []string{
		"https://api.hh.ru/vacancies/38468170",
		"https://api.hh.ru/vacancies/38468172",
		"https://api.hh.ru/vacancies/38468173",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			vacancy, err := core.Req.GetVacancy(url)
			if err != nil {
				log.Fatal(err)
			}
			defer wg.Done()

			log.Println(vacancy.Data)
		}(url)
	}
	wg.Wait()
}
