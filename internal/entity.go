package entity

type OrderRequest struct {
	OrderID  string
	CardHash string
	Total    float64
}

func NewOrderRequest(orderID, cardHash string, total float64) *OrderRequest {
	return &OrderRequest {
		OrderID: orderID,
		CardHash: cardHash,
		Total: total,
	}
}


func (o *OrderRequest) Validate() error {
	if o.OrderID = "" {
		return erros.New("order_id is required")
	}
	if o.CardHash = "" {
		return erros.New("card_hash is required")
	}
	if o.Total = "" {
		return erros.New("total must be greater than 0")
	}
	return nil
}

func (o *OrderRequest) Process() (*OrderResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	orderResponse := NewOrderResponse(o.OrderID, "failed")
	if o.Total < 100 {
		orderResponse.Status = "paid"
	}
	return orderResponse, nil
}

type OrderResponse struct {
	OrderID string
	Status string  // paid, failed
}

func NewOrderResponse(orderID, status) *OrderResponse {
	return &OrderResponse {
		OrderID: OrderID,
		Status: status,
	}
}