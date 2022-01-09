package service_test

import (
	"database/sql"
	"github.com/golang/mock/gomock"
	"net/http"
	"test/errs"
	"test/mock/mock_repository"
	"test/service"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)

	customerID := 1

	mockRepo.EXPECT().GetById(customerID).Return(nil, sql.ErrNoRows)
	customerService := service.NewCustomerService(mockRepo)

	_, err := customerService.GetCustomer(customerID)

	if err == nil {
		t.Error("should be error")
		return
	}

	appErr, ok := err.(errs.AppError)
	if !ok {
		t.Error("should return AppError")
		return
	}

	if appErr.Code != http.StatusNotFound {
		t.Error("invalid error code")
	}

	if appErr.Message != "customer not found" {
		t.Error("invalid error message")
	}
}
