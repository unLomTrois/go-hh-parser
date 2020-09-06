package main

import (
	"go-hh-parser/internal/app/core"
	"log"
	"sync"
)

func main() {
	Core := core.New()

	urls := []string{
		"https://api.hh.ru/vacancies/38461984",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			vacancy, err := Core.Requests.GetVacancy(url)
			if err != nil {
				log.Fatal(err)
			}
			defer wg.Done()

			// get keyskills
			keyskills := (*vacancy.Data)["key_skills"].(core.Interarr)
			for _, skill := range keyskills {
				log.Println(skill.(core.Intermap)["name"])
			}

			log.Println(vacancy.Data)
		}(url)
	}
	wg.Wait()
}
