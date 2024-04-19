package service

import (
	"encoding/json"
	"errors"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/vo"
)

var FinanceLogServiceImpl FinanceLogService

type FinanceLogService struct {
	SimpleService[model.FinanceLog, uint64]
}

func (s *FinanceLogService) UpsertFinance(financeLog *vo.FinanceLog) error {
	if financeLog == nil {
		return errors.New("参数不合法")
	}

	bs, _ := json.Marshal(&model.FinanceLogContent{
		Savings:     financeLog.Savings,
		Income:      financeLog.Income,
		Expenditure: financeLog.Expenditure,
	})
	return s.Upsert(&model.FinanceLog{
		Month:   financeLog.Month,
		Content: string(bs),
		Total:   financeLog.Total,
	})
}

func (s *FinanceLogService) GetFinance(month string) (*vo.FinanceLog, error) {
	financeLog, err := dao.FinanceLogDaoImpl.GetByMonth(month)
	if err != nil {
		return nil, err
	}
	if financeLog == nil {
		return nil, nil
	}

	content := &model.FinanceLogContent{}
	err = json.Unmarshal([]byte(financeLog.Content), content)
	if err != nil {
		return nil, err
	}

	return &vo.FinanceLog{
		Month:       financeLog.Month,
		Savings:     content.Savings,
		Income:      content.Income,
		Expenditure: content.Expenditure,
		Total:       financeLog.Total,
	}, nil
}
