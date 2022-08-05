package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Sample struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetSample(id int) (sample Sample, err error) {
	sample = Sample{}
	err = DB.QueryRow(`
		SELECT id, title, content, photo
			FROM "samples"
			WHERE "id"=$1
	`, id).
		Scan(&sample.Id, &sample.Title, &sample.Content, &sample.Photo)

	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			fmt.Println(`No sample with id:`, id)
		case err != nil:
			LogMessage("Error when try to GetSample with id: " + strconv.Itoa(id))
		}
	}
	return
}

func GetSamples() (samples []Sample, err error) {
	samples = make([]Sample, 0)

	rows, queryErr := DB.Query(`
		SELECT id, title, content, photo
			FROM "samples"
			LIMIT 20
	`)

	if queryErr != nil {
		switch {
		case queryErr == sql.ErrNoRows:
			fmt.Println("Samples is empty")
		case queryErr != nil:
			LogMessage("Error when try to get all samples")
		}
		err = queryErr
	}

	defer rows.Close()

	for rows.Next() {
		sample := Sample{}
		if err := rows.Scan(&sample.Id, &sample.Title, &sample.Content, &sample.Photo); err != nil {
			log.Fatal(err)
		}
		samples = append(samples, sample)
	}

	return
}

func (sample *Sample) Create() (err error) {
	statement := `INSERT INTO "samples"
		("title", "content", "photo", "created_at", "updated_at") 
		values ($1, $2, $3, now(), now()) returning id`
	stmt, err := DB.Prepare(statement)

	defer stmt.Close()
	err = stmt.QueryRow(sample.Title, sample.Content, sample.Photo).
		Scan(&sample.Id)

	if err != nil && err == sql.ErrNoRows {
		fmt.Println(`Cannot create sample with title:`, sample.Title)
		msg := "Cannot create sample with title: " + sample.Title
		LogMessage("Cannot create sample with info: " + msg)
	}

	return
}

// Consider very carefully when use this method. It will remove record in database
func (sample *Sample) HardDelete() (res sql.Result, err error) {
	res, err = DB.Exec(`
		DELETE FROM "Samples"
		WHERE "id" = $1
	`, sample.Id)

	return
}
