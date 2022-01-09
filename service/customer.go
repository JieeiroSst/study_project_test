package service

import (
	"database/sql"
	"test/errs"
	"test/model"
	"test/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

//go:generate mockgen -destination=../mock/mock_service/mock_customer_service.go test/service CustomerService
type CustomerService interface {
	GetCustomers() ([]model.CustomerResponse, error)
	GetCustomer(int) (*model.CustomerResponse, error)
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return &customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]model.CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	custResponses := []model.CustomerResponse{}
	for _, customer := range customers {
		custResponse := model.CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*model.CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewValidationError("customer not found")
		}

		return nil, errs.NewUnexpectedError()
	}

	custResponse := model.CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
