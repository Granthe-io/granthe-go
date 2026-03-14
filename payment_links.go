package flow

import "context"

type PaymentLinkService struct{ http *httpClient }

type CreatePaymentLinkParams struct {
	Title          string  `json:"title"`
	Description    string  `json:"description,omitempty"`
	Amount         float64 `json:"amount,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	CryptoCurrency string  `json:"crypto_currency,omitempty"`
	Network        string  `json:"network,omitempty"`
	RedirectURL    string  `json:"redirect_url,omitempty"`
}

func (s *PaymentLinkService) Create(ctx context.Context, params *CreatePaymentLinkParams) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.post(ctx, "/merchants/me/payment-links", params, &result)
	return result, err
}

func (s *PaymentLinkService) List(ctx context.Context, params map[string]string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.http.get(ctx, "/merchants/me/payment-links", params, &result)
	return result, err
}

func (s *PaymentLinkService) Update(ctx context.Context, linkID string, params map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.patch(ctx, "/merchants/me/payment-links/"+linkID, params, &result)
	return result, err
}

func (s *PaymentLinkService) Delete(ctx context.Context, linkID string) error {
	return s.http.delete(ctx, "/merchants/me/payment-links/"+linkID)
}
