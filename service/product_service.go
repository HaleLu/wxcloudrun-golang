package service

import (
	"wxcloudrun-golang/vo"

	"wxcloudrun-golang/db/model"
)

var ProductServiceImpl ProductService

type ProductService struct {
	SimpleService[model.Product, uint64]
}

// UpsertProduct 创建或修改商品
func (s *ProductService) UpsertProduct(req *vo.ProductReq) error {
	return s.Upsert(req.Product)
}

// DeleteProduct 删除商品
func (s *ProductService) DeleteProduct(id int64) error {
	return s.Delete(uint64(id))
}

// GetProduct 查询当前商品
func (s *ProductService) GetProduct(id int64) (*model.Product, error) {
	return s.Get(uint64(id))
}

// ListProduct 查询当前商品
func (s *ProductService) ListProduct() ([]*model.Product, error) {
	return s.List()
}
