package entities

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreditDetails es el holder donde se lleva la información del Credito que se requiere calcular
type CreditDetails struct {
	Count        int32 `bson:"count"`
	LoanQuantity int32 `bson:"loan_quantity"`
}

func (c CreditDetails) ToString() string {
	return fmt.Sprintf("CreditDetail{LoanQuantity:%d, Count: %d, TotalWithCount: %d}\n", c.LoanQuantity, c.Count, c.TotalWithCount())
}

func (c CreditDetails) TotalWithCount() int32 {
	return c.LoanQuantity * c.Count
}

type CreditDetailsWithStatus struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Investment     int32              `bson:"investment"`
	CreditsDetails []*CreditDetails   `bson:"credits_details"`
	Status         string             `bson:"status"`
}
