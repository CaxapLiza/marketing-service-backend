package internal

type Analytic struct {
	ID               int     `json:"id"`
	CTR              float32 `json:"ctr"`
	ConversationRate float32 `json:"conversation_rate"`
	ROI              float32 `json:"roi"`
	CAC              float32 `json:"cac"`
	LTV              float32 `json:"ltv"`
	ProjectID        int     `json:"project_id"`
}

type AnalyticInputData struct {
	NumberOfImpressions             int     `json:"number_of_impressions"`
	NumberOfClicks                  int     `json:"number_of_clicks"`
	NumberOfVisitors                int     `json:"number_of_visitors"`
	NumberOfConversions             int     `json:"number_of_conversions"`
	Profit                          float32 `json:"profit"`
	Costs                           float32 `json:"costs"`
	NumberOfNewCustomers            int     `json:"number_of_new_customers"`
	MarketingCosts                  int     `json:"marketing_costs"`
	AverageAnnualRevenuePerCustomer float32 `json:"average_annual_revenue_per_customer"`
	AverageCustomerLifespan         float32 `json:"average_customer_lifespan"`
	ProjectID                       int     `json:"project_id"`
}
