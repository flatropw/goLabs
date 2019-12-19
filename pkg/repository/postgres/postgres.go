package postgres

import (
	"database/sql"
	"fmt"
	dberror "github.com/Shyp/go-dberror"
	"github.com/burov/task_databases/pkg/model"
)

const connString = "postgres://ehcqpgdd:gQ2UB2W-qpFBKPKEU9aDyLqTyWwLs4ED@manny.db.elephantsql.com:5432/ehcqpgdd"

func Connect() (*sql.DB, error) {
	return sql.Open("postgres", connString)
}

type ContactsRepositoryPostgres struct {
	db *sql.DB
}

func NewContactsRepositoryPostgres(db *sql.DB) *ContactsRepositoryPostgres {
	return &ContactsRepositoryPostgres{db: db}
}

func (r *ContactsRepositoryPostgres) Save(contact model.Contact) (model.Contact, error) {
	err := r.db.QueryRow(InsertQuery, contact.FirstName, contact.LastName, contact.Phone, contact.Email).Scan(&contact.ID)
	dbErr := dberror.GetError(err)
	switch e := dbErr.(type) {
	case *dberror.Error:
		return model.Contact{}, fmt.Errorf(e.Error())
	default:
		return contact, nil
	}
}

func (r *ContactsRepositoryPostgres) ListAll() (contacts []model.Contact, err error) {
	rows, err := r.db.Query(ListAllQuery)
	defer func() {
		_ = rows.Close()
	}()
	dbErr := dberror.GetError(err)
	switch e := dbErr.(type) {
	case *dberror.Error:
		return contacts, fmt.Errorf(e.Error())
	default:
		for rows.Next() {
			var contact model.Contact
			err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
			if err != nil {
				return
			}
			contacts = append(contacts, contact)
		}
	}

	return
}

func (r *ContactsRepositoryPostgres) GetByID(id uint) (contact model.Contact, err error) {
	row := r.db.QueryRow(GetByIdQuery, id)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
	dbErr := dberror.GetError(err)
	switch e := dbErr.(type) {
	case *dberror.Error:
		return model.Contact{}, fmt.Errorf(e.Error())
	default:
		return
	}
}

func (r *ContactsRepositoryPostgres) GetByPhone(phone string) (contact model.Contact, err error) {
	row := r.db.QueryRow(GetByPhoneQuery, phone)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
	dbErr := dberror.GetError(err)
	switch e := dbErr.(type) {
	case *dberror.Error:
		return model.Contact{}, fmt.Errorf(e.Error())
	default:
		return
	}
}

func (r *ContactsRepositoryPostgres) GetByEmail(email string) (contact model.Contact, err error) {
	row := r.db.QueryRow(GetByEmailQuery, email)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
	dbErr := dberror.GetError(err)
	switch e := dbErr.(type) {
	case *dberror.Error:
		return model.Contact{}, fmt.Errorf(e.Error())
	default:
		return
	}
}

func (r *ContactsRepositoryPostgres) SearchByName(n string) (contacts []model.Contact, err error) {
	rows, err := r.db.Query(SearchByNameQuery, n)
	defer func() {
		_ = rows.Close()
	}()
	dbErr := dberror.GetError(err)
	switch e := dbErr.(type) {
	case *dberror.Error:
		return contacts, fmt.Errorf(e.Error())
	}

	for rows.Next() {
		var contact model.Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		dbErr := dberror.GetError(err)
		switch e := dbErr.(type) {
		case *dberror.Error:
			return contacts, fmt.Errorf(e.Error())
		default:
			contacts = append(contacts, contact)
		}
	}

	return
}

func (r *ContactsRepositoryPostgres) Delete(id uint) error {
	_, err := r.db.Exec(DeleteQuery, id)
	return dberror.GetError(err)
}
