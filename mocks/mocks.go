package mocks

import (
	"github.com/manicar2093/YoFioExamen/entities"
	"github.com/manicar2093/YoFioExamen/models"
	"github.com/stretchr/testify/mock"
)

// Services Mocks
type InvestmentFilter struct {
	mock.Mock
}

func (c InvestmentFilter) Assign(investment int32) (int32, int32, int32, error) {
	args := c.Called(investment)
	return args.Get(0).(int32), args.Get(1).(int32), args.Get(2).(int32), args.Error(3)

}

type InvestmentFilterMock struct {
	mock.Mock
}

func (i InvestmentFilterMock) Filter(quantity int32, credit1, credit2, credit3 *entities.CreditDetails) (e error) {
	args := i.Called(quantity, credit1, credit2, credit3)
	return args.Error(0)
}

type CreditDetailsServiceMock struct {
	mock.Mock
}

func (c CreditDetailsServiceMock) GetAllCreditDetails() ([]entities.CreditDetails, error) {
	args := c.Called()
	return args.Get(0).([]entities.CreditDetails), args.Error(1)
}

func (c CreditDetailsServiceMock) SaveSuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error {
	args := c.Called(credit1, credit2, credit3, invest)
	return args.Error(0)
}

func (c CreditDetailsServiceMock) SaveUnsuccessfulRequest(credit1 *entities.CreditDetails, credit2 *entities.CreditDetails, credit3 *entities.CreditDetails, invest int32) error {
	args := c.Called(credit1, credit2, credit3, invest)
	return args.Error(0)
}

func (c CreditDetailsServiceMock) GetStatistics() (models.CreditsAssignmentStatistics, error) {
	args := c.Called()
	return args.Get(0).(models.CreditsAssignmentStatistics), args.Error(1)
}

// DAO Mocks

type CreditDetailsDaoMock struct {
	mock.Mock
}

func (c CreditDetailsDaoMock) FilterCreditDetailsWithStatus(filter interface{}) ([]entities.CreditDetailsWithStatus, error) {
	args := c.Called(filter)
	return args.Get(0).([]entities.CreditDetailsWithStatus), args.Error(1)
}

func (c CreditDetailsDaoMock) Save(data *entities.CreditDetailsWithStatus) error {
	args := c.Called(data)
	return args.Error(0)
}
