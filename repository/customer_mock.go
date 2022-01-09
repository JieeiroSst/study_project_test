package repository

import (
	"errors"
	"test/model"
)

type customerRepositoryMock struct {
	customers []model.Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []model.Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 1},
		{CustomerID: 1002, Name: "Rob", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 0},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]model.Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*model.Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}

