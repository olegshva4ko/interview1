package db_test_test

import (
	"encoding/json"
	"interview/pkg/database"
	"testing"
)

//TestWriteDB tests if the writing to db is correct
func TestWriteDB(t *testing.T) {
	sqlite, err := database.MakeSQlite("test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer sqlite.DB.Close()
	defer sqlite.DB.Exec("DROP TABLE Test")
	type user struct {
		ID      int `json:"-"`
		Message string
		Name    string
	}

	messages := [4]string{
		"{\"message\": \"test1\"}",
		"{\"message\": \"test2\", \"name\": \"mike\"}",
		"{\"name\": \"Michael\"}",
		"{\"message\": \"test4\", \"name\": \"Misha\"}",
	}

	for i := range messages {
		u := new(user)
		err := json.Unmarshal([]byte(messages[i]), &u)
		if err != nil {
			t.Fatal(err)
		}
		sqlite.TestHandler(u.Message, u.Name)
	}

	//read from sqlite
	rows, err := sqlite.DB.Query("SELECT * FROM Test")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	results := []*user{}
	for rows.Next() {
		result := new(user)
		err := rows.Scan(
			&result.ID,
			&result.Message,
			&result.Name,
		)
		if err != nil {
			t.Error(err)
		}
		results = append(results, result)
	}

	expectedResult := []*user{
		{
			Message: "test2",
			Name:    "mike",
		},
		{
			Message: "test4",
			Name:    "Misha",
		},
	}
	if len(results) != len(expectedResult) {
		t.FailNow()
	}
	for i := range results {
		if expectedResult[i].Message != results[i].Message {
			t.Errorf("%s %s", expectedResult[i].Message, results[i].Message)
		}
		if expectedResult[i].Name != results[i].Name {
			t.Errorf("%s %s", expectedResult[i].Name, results[i].Name)
		}

	}

}
