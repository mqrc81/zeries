// Package cmd/jobs is the entry-point for a daily jobs
// which notifies users of new releases & recommendations
// and updates list of upcoming releases stored in db
package main

import (
	"os"

	"github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
	"github.com/mqrc81/zeries/jobs"
	"github.com/mqrc81/zeries/postgres"
	"github.com/mqrc81/zeries/trakt"
	. "github.com/mqrc81/zeries/util"
)

func main() {
	// Environment variables need to be initialized from .env file first when ran locally
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		err := godotenv.Load()
		checkError(err)
	}

	store, _, err := postgres.NewStore(os.Getenv("DATABASE_URL"))
	checkError(err)

	tmdbClient, err := tmdb.Init(os.Getenv("TMDB_KEY"))
	checkError(err)

	traktClient, err := trakt.Init(os.Getenv("TRAKT_KEY"))
	checkError(err)

	err = jobs.NewUpdateReleasesJob(*store, tmdbClient, traktClient).Execute()
	checkError(err)

	err = jobs.NewNotifyUsersJob().Execute()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		LogPanic(err)
	}
}
