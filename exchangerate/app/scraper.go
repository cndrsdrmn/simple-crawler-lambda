package app

import (
	_ "exchangerate/config"
	_ "exchangerate/database"
	repo "exchangerate/repositories/exchangerate"

	"exchangerate/spiders"
)

func RunScraper() error {

	scraping := spiders.Exec()

	if err := repo.New().Upsert(scraping); err != nil {
		return err
	}

	return nil
}
