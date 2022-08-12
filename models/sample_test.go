package models

import (
	"database/sql"
	"net/http"
	"testing"
)

const TestId = "6fe370c7-0807-43eb-9d45-dc540aa6cc73"
const TestTitle = "Sample 1"
const TestContent = "Lorem ipsum content 1"
const TestPhoto = ""

var TestSample = Sample{
	Title:   "create to delete test sample",
	Content: "create to delete test sample",
	Photo:   "create-to-delete-test-sample",
}

func TestGetSample(t *testing.T) {
	sample, err := GetSample(TestId)

	if err != nil {
		t.Log(err)
		t.Errorf("Cannot get sample with id: %s", TestId)
	}
	if sample.Id != TestId {
		t.Errorf("Sample id is is different with the expected TestId")
	}
}

func TestCreateThenDeleteSample(t *testing.T) {
	id, createErr := TestSample.Create()
	if createErr != nil {
		t.Errorf("Cannot create Sample!!!")
	}

	deleteRes, deleteErr := HardDelete(id)
	status, err := checkEffected(deleteRes)

	if status != "" {
		t.Errorf(status)
	}

	if err != nil || deleteErr != nil {
		t.Errorf("Cannot delete sample!!!")
	}

}

func checkEffected(res sql.Result) (status string, err error) {
	rowEffect, err := res.RowsAffected()

	if rowEffect == 0 {
		status = http.StatusText(400)
	}

	return
}
