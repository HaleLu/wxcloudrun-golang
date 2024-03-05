package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wxcloudrun-golang/vo"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// getAction 获取action
func getProductReq(r *http.Request) (*vo.ProductReq, error) {
	decoder := json.NewDecoder(r.Body)
	body := &vo.ProductReq{}
	if err := decoder.Decode(body); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if body.Action == "" || body.Product == nil {
		return nil, fmt.Errorf("缺少参数")
	}

	return body, nil
}

// ProductHandler 商品接口
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	fmt.Printf("req:%+v\n", r)
	if r.Method == http.MethodGet {
		id, _ := strconv.ParseInt(r.Form.Get("id"), 10, 64)
		product, err := getProduct(id)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = product
		}
	} else if r.Method == http.MethodPost {
		product, err := modifyProduct(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = product
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// modifyCounter 更新计数，自增或者清零
func modifyProduct(r *http.Request) (*model.Product, error) {
	req, err := getProductReq(r)
	if err != nil {
		return nil, err
	}

	if req.Action == "inc" {
		req.Product, err = upsertProduct(req)
		if err != nil {
			return nil, err
		}
	} else if req.Action == "clear" {
		err = deleteProduct(int64(req.Product.Id))
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", req.Action)
	}

	return req.Product, err
}

// upsertCounter 更新或修改计数器
func upsertProduct(req *vo.ProductReq) (*model.Product, error) {
	err := dao.ProductDaoImpl.Upsert(req.Product)
	if err != nil {
		return nil, err
	}
	return req.Product, nil
}

func deleteProduct(id int64) error {
	return dao.ProductDaoImpl.Delete(id)
}

// getProduct 查询当前商品
func getProduct(id int64) (*model.Product, error) {
	product, err := dao.ProductDaoImpl.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
