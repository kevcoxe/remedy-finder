package db

type Remedy struct {
	Id          int64  `db:Id`
	Name        string `db:Name`
	Description string `db:Description`
}

type Symptom struct {
	Id          int64  `db:Id`
	Name        string `db:Name`
	Description string `db:Description`
}

type DBClient interface {
	GetRemedies() []Remedy
	GetRemedyByName(name string) (*Remedy, error)
	GetRemedyById(id string) (*Remedy, error)
	CreateRemedy(Remedy) (*Remedy, error)
	UpdateRemedyById(id string, r Remedy) (Remedy, error)

	GetSymptoms() []Symptom
	GetSymptomByName(name string) (*Symptom, error)
	GetSymptomById(id string) (*Symptom, error)
	CreateSymptom(Symptom) (*Symptom, error)
	UpdateSymptomById(id string, s Symptom) (*Symptom, error)
}
