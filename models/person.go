// For now, if the is file changed, it must be copied to
// C:\Program Files\Go\src\models
// BEFORE issuing 'go build' command
// Or Use 'make' in the PersonWeb folder which copies all files in the 'models' folder whether they changed or not
//   - You can run make.bat via Command Prompt, but Command Prompt must be RUN AS ADMINISTRATOR
//   - You can doulbe click the 'make' shortcut

package models

import (
	"database/sql"
	"strconv"
  _ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite", "./names.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Person struct {
	Id         int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email string   `json:"email"`
	IpAddress string `json:"ip_address"`
}

func GetPersons(count int) ([]Person, error) {

	rows, err := DB.Query("SELECT id, first_name, last_name, email, ip_address from people LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	people := make([]Person, 0)

	for rows.Next() {
		singlePerson := Person{}
    err = rows.Scan(&singlePerson.Id, &singlePerson.FirstName, &singlePerson.LastName, &singlePerson.Email, &singlePerson.IpAddress)

		if err != nil {
			return nil, err
		}

		people = append(people, singlePerson)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return people, err
}