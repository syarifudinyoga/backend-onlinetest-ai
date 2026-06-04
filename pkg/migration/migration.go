package migration

import (
	// "context"
	"context"
	"io/ioutil"
	"log"
	"online-test/config"

	// "os"
	"sort"
	"strings"
	// "online-test/config"
)

func RunMigrations(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("failed to read migrations folder:", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		if isExecuted(file.Name()) {
			continue
		}

		runFile(path + "/" + file.Name())
		markExecuted(file.Name())
	}
}

func markExecuted(filename string) {
	_, err := config.DB.Exec(context.Background(),
		`INSERT INTO schema_migrations (filename) VALUES ($1)`,
		filename,
	)

	if err != nil {
		log.Fatal("failed marking migration:", filename, err)
	}
}

func runFile(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("failed reading file:", filePath, err)
	}

	_, err = config.DB.Exec(context.Background(), string(content))
	if err != nil {
		log.Fatal("failed executing migration:", filePath, err)
	}

	log.Println("migration executed:", filePath)
}

func isExecuted(filename string) bool {
	var exists bool

	err := config.DB.QueryRow(context.Background(),
		`SELECT EXISTS (
			SELECT 1 FROM schema_migrations WHERE filename=$1
		)`, filename,
	).Scan(&exists)

	if err != nil {
		log.Fatal(err)
	}

	return exists
}
