package repository

import (
	"github.com/jmoiron/sqlx"
	"test/model"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

//go:generate mockgen -destination=../mock/mock_repository/mock_customer_repository.go test/repository CustomerRepository
type CustomerRepository interface {
	GetAll() ([]model.Customer, error)
	GetById(int) (*model.Customer, error)
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return &customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]model.Customer, error) {
	customers := []model.Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*model.Customer, error) {
	customer := model.Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
