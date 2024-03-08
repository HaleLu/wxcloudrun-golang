package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/service"
	"wxcloudrun-golang/vo"
)

// ProductsHandler 商品列表接口
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	fmt.Printf("req:%+v\n", r)
	if r.Method == http.MethodGet {
		products, err := service.ProductServiceImpl.ListProduct()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = products
		}
	}
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// ProductHandler 商品接口
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	openId := GetOpenId(r)
	fmt.Println("open_id: " + openId)

	res := &JsonResult{}
	fmt.Printf("req:%+v\n", r)
	if r.Method == http.MethodGet {
		id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		product, err := service.ProductServiceImpl.GetProduct(id)
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
	_, _ = w.Write(msg)
}

// modifyProduct 更新商品
func modifyProduct(r *http.Request) (*model.Product, error) {
	req, err := getReq[vo.ProductReq](r)
	if err != nil {
		return nil, err
	}
	if req.Action == "" || req.Product == nil {
		return nil, fmt.Errorf("缺少参数")
	}

	if req.Action == "inc" {
		err = service.ProductServiceImpl.UpsertProduct(req)
		if err != nil {
			return nil, err
		}
	} else if req.Action == "clear" {
		err = service.ProductServiceImpl.DeleteProduct(int64(req.Product.Id))
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", req.Action)
	}

	return req.Product, err
}
