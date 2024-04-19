package vo

import "wxcloudrun-golang/db/model"

type FinanceLogReq struct {
	Action     string     `json:"action"`
	FinanceLog FinanceLog `json:"finance_log"`
}

type FinanceLog struct {
	Month       string        `json:"month"`
	Savings     []*model.Item `json:"savings"`
	Income      []*model.Item `json:"income"`
	Expenditure []*model.Item `json:"expenditure"`
	Total       int64         `json:"total"`
}
