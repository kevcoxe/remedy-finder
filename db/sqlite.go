package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteClient struct {
	db *sql.DB
}

func (sc *SqliteClient) GetRemedies() []Remedy {
	// Execute the query
	rows, err := sc.db.Query("SELECT * FROM remedies")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	remedies := []Remedy{}
	// Iterate over the rows and scan them into your struct
	for rows.Next() {
		var remedy Remedy
		err := rows.Scan(&remedy.Id, &remedy.Name, &remedy.Description)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%+v\n", remedy)
		remedies = append(remedies, remedy)
	}

	return remedies
}

func (sc *SqliteClient) GetRemedyByName(name string) (*Remedy, error) {
	// Execute the query
	rows, err := sc.db.Query(fmt.Sprintf("SELECT * FROM remedies WHERE Name = %v", name))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	remedy := Remedy{}

	// Iterate over the rows and scan them into your struct
	if !rows.Next() {
		return nil, errors.New("no remedies found")
	}

	remedyErr := rows.Scan(&remedy.Id, &remedy.Name, &remedy.Description)
	if remedyErr != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", remedy)

	return &remedy, nil
}

func (sc *SqliteClient) GetRemedyById(id string) (*Remedy, error) {
	// Execute the query
	rows, err := sc.db.Query(fmt.Sprintf("SELECT * FROM remedies WHERE Id = %v", id))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	remedy := Remedy{}

	// Iterate over the rows and scan them into your struct
	if !rows.Next() {
		return nil, errors.New("no remedies found")
	}

	remedyErr := rows.Scan(&remedy.Id, &remedy.Name, &remedy.Description)
	if remedyErr != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", remedy)

	return &remedy, nil
}

func (sc *SqliteClient) CreateRemedy(name string, description string) (*Remedy, error) {
	// Prepare the SQL statement
	stmt, err := sc.db.Prepare(`
		INSERT INTO remedies (Name, Description)
		VALUES (?,?)
		RETURNING Id, Name, Description;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	r := &Remedy{}

	// Execute the prepared statement
	err = stmt.QueryRow(name, description).Scan(&r.Id, &r.Name, &r.Description)
	if err != nil {
		fmt.Printf("unable to insert data: %v", err)
		return nil, err
	}

	return r, nil
}

func (sc *SqliteClient) UpdateRemedyById(id string, name string, description string) (*Remedy, error) {
	r := &Remedy{}
	return r, nil
}

func (sc *SqliteClient) GetSymptoms() []Symptom {
	// Execute the query
	rows, err := sc.db.Query("SELECT * FROM symptoms")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	symptoms := []Symptom{}
	// Iterate over the rows and scan them into your struct
	for rows.Next() {
		var symptom Symptom
		err := rows.Scan(&symptom.Id, &symptom.Name, &symptom.Description)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%+v\n", symptom)
		symptoms = append(symptoms, symptom)
	}

	return symptoms
}

func (sc *SqliteClient) GetSymptomByName(name string) (*Symptom, error) {
	// Execute the query
	rows, err := sc.db.Query(fmt.Sprintf("SELECT * FROM symptoms WHERE Name = %v", name))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	symptom := Symptom{}

	// Iterate over the rows and scan them into your struct
	if !rows.Next() {
		return nil, errors.New("no remedies found")
	}

	remedyErr := rows.Scan(&symptom.Id, &symptom.Name, &symptom.Description)
	if remedyErr != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", symptom)

	return &symptom, nil
}

func (sc *SqliteClient) GetSymptomById(id string) (*Symptom, error) {
	// Execute the query
	rows, err := sc.db.Query(fmt.Sprintf("SELECT * FROM symptoms WHERE Id = %v", id))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	symptom := Symptom{}

	// Iterate over the rows and scan them into your struct
	if !rows.Next() {
		return nil, errors.New("no remedies found")
	}

	remedyErr := rows.Scan(&symptom.Id, &symptom.Name, &symptom.Description)
	if remedyErr != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", symptom)

	return &symptom, nil
}

func (sc *SqliteClient) CreateSymptom(name string, description string) (*Symptom, error) {
	// Prepare the SQL statement
	stmt, err := sc.db.Prepare(`
		INSERT INTO symptoms (Name, Description)
		VALUES (?,?)
		RETURNING Id, Name, Description;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	r := &Symptom{}

	// Execute the prepared statement
	err = stmt.QueryRow(name, description).Scan(&r.Id, &r.Name, &r.Description)
	if err != nil {
		fmt.Printf("unable to insert data: %v", err)
		return nil, err
	}

	return r, nil
}

func (sc *SqliteClient) UpdateSymptomById(id string, s *Symptom) (Symptom, error) {
	return *s, nil
}

func (sc *SqliteClient) InitiateTable() {
	// Ensure the table exists
	_, err := sc.db.Exec(`
	CREATE TABLE IF NOT EXISTS remedies (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL UNIQUE,
			Description TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS symptoms (
			Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL UNIQUE,
			Description TEXT NOT NULL
	);

	INSERT or IGNORE INTO remedies (Id, Name, Description)
	VALUES
	(0, 'Vitamin A', 'A vitamin'),
	(1, 'Vitamin B', 'A vitamin'),
	(2, 'Vitamin C', 'A vitamin');

	`)
	if err != nil {
		panic(err.Error())
	}
}

func (sc *SqliteClient) CloseDB() {
	sc.db.Close()
}

func (sc *SqliteClient) Ping() error {
	err := sc.db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Connected!")
	return nil
}

func NewSQLiteClient() (SqliteClient, error) {

	db, err := sql.Open("sqlite3", "./myDBname.db")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected!")

	client := SqliteClient{
		db: db,
	}

	client.InitiateTable()

	return client, nil

}
