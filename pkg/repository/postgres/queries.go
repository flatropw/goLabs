package postgres

const (
	InsertQuery = "INSERT INTO contacts (first_name, last_name, phone, email) VALUES ($1,$2,$3,$4);"
	ListAllQuery = "SELECT id, first_name, last_name, phone, email FROM contacts;"
	GetByIdQuery = "SELECT id, first_name, last_name, phone, email FROM contacts WHERE id = $1 LIMIT 1;"
	GetByPhoneQuery = "SELECT id, first_name, last_name, phone, email FROM contacts WHERE phone = $1 LIMIT 1;"
	GetByEmailQuery = "SELECT id, first_name, last_name, phone, email FROM contacts WHERE email = $1 LIMIT 1;"
	SearchByNameQuery = "SELECT id, first_name, last_name, phone, email FROM contacts WHERE first_name = $1;"
	DeleteQuery = "DELETE FROM contacts WHERE id = $1;"
)
