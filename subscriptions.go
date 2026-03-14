package flow

import "context"

type SubscriptionService struct{ http *httpClient }

type CreatePlanParams struct {
	Name           string  `json:"name"`
	Amount         float64 `json:"amount"`
	Description    string  `json:"description,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	IntervalType   string  `json:"interval_type,omitempty"`
	IntervalCount  int     `json:"interval_count,omitempty"`
	TrialDays      int     `json:"trial_days,omitempty"`
	MaxSubscribers int     `json:"max_subscribers,omitempty"`
	WebhookURL     string  `json:"webhook_url,omitempty"`
}

func (s *SubscriptionService) CreatePlan(ctx context.Context, params *CreatePlanParams) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.post(ctx, "/merchants/me/subscription-plans", params, &result)
	return result, err
}

func (s *SubscriptionService) ListPlans(ctx context.Context, params map[string]string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.http.get(ctx, "/merchants/me/subscription-plans", params, &result)
	return result, err
}

func (s *SubscriptionService) UpdatePlan(ctx context.Context, planID string, params map[string]interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.patch(ctx, "/merchants/me/subscription-plans/"+planID, params, &result)
	return result, err
}

func (s *SubscriptionService) DeletePlan(ctx context.Context, planID string) error {
	return s.http.delete(ctx, "/merchants/me/subscription-plans/" + planID)
}

func (s *SubscriptionService) ListSubscriptions(ctx context.Context, params map[string]string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.http.get(ctx, "/merchants/me/subscriptions", params, &result)
	return result, err
}

func (s *SubscriptionService) CancelSubscription(ctx context.Context, subID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.post(ctx, "/merchants/me/subscriptions/"+subID+"/cancel", nil, &result)
	return result, err
}

func (s *SubscriptionService) GetStats(ctx context.Context) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.get(ctx, "/merchants/me/subscriptions/stats", nil, &result)
	return result, err
}

func (s *SubscriptionService) ListPayments(ctx context.Context, params map[string]string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := s.http.get(ctx, "/merchants/me/subscription-payments", params, &result)
	return result, err
}
