package main

import (
	"go-hh-parser/internal/app/core"
	"log"
)

func main() {
	Core := core.New()

	_, err := Core.Requests.SearchVacancies(
		core.VacancyQueryParams{
			Text:     "разработчик",
			Area:     "1",
			NoMagic:  true,
			Page:     0,
			PerPage:  100,
			OrderBy:  "salary_desc",
			Clusters: true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// jsonData, err := json.MarshalIndent(&(*pages.Clusters)[4].Items, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(string(jsonData))
	// log.Println((*pages.Clusters)[4].Items)
	// log.Println(len(pages.Itemsх))

	// urls := []string{
	// 	"https://api.hh.ru/vacancies/38840984",
	// }

	// var wg sync.WaitGroup
	// for _, url := range urls {
	// 	wg.Add(1)

	// 	go func(url string) {
	// 		vacancy, err := Core.Requests.GetFullVacancy(url)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		defer wg.Done()

	// 		// get keyskills
	// 		for _, skill := range vacancy.Keyskills {
	// 			log.Println(skill["name"])
	// 		}

	// 		jsonData, err := json.MarshalIndent(&vacancy, "", "  ")
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		log.Println(string(jsonData))
	// 	}(url)
	// }
	// wg.Wait()
}
