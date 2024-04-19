package handler

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"wxcloudrun-golang/db/model"
//	"wxcloudrun-golang/service"
//	"wxcloudrun-golang/vo"
//)
//
//// FinanceLogsHandler 财务记录列表接口
//func FinanceLogsHandler(w http.ResponseWriter, r *http.Request) {
//	res := &JsonResult{}
//	fmt.Printf("req:%+v\n", r)
//	if r.Method == http.MethodGet {
//		financeLogs, err := service.FinanceLogServiceImpl.List()
//		if err != nil {
//			res.Code = -1
//			res.ErrorMsg = err.Error()
//		} else {
//			res.Data = financeLogs
//		}
//	}
//	msg, err := json.Marshal(res)
//	if err != nil {
//		fmt.Fprint(w, "内部错误")
//		return
//	}
//	w.Header().Set("content-type", "application/json")
//	w.Write(msg)
//}
//
//// FinanceLogHandler 商品接口
//func FinanceLogHandler(w http.ResponseWriter, r *http.Request) {
//	openId := GetOpenId(r)
//	fmt.Println("open_id: " + openId)
//
//	res := &JsonResult{}
//	fmt.Printf("req:%+v\n", r)
//	if r.Method == http.MethodGet {
//		month := r.URL.Query().Get("month")
//		product, err := service.FinanceLogServiceImpl.GetFinance(month)
//		if err != nil {
//			res.Code = -1
//			res.ErrorMsg = err.Error()
//		} else {
//			res.Data = product
//		}
//	} else if r.Method == http.MethodPost {
//		product, err := modifyProduct(r)
//		if err != nil {
//			res.Code = -1
//			res.ErrorMsg = err.Error()
//		} else {
//			res.Data = product
//		}
//	} else {
//		res.Code = -1
//		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
//	}
//
//	msg, err := json.Marshal(res)
//	if err != nil {
//		fmt.Fprint(w, "内部错误")
//		return
//	}
//	w.Header().Set("content-type", "application/json")
//	_, _ = w.Write(msg)
//}
//
//// modifyProduct 更新商品
//func modifyProduct(r *http.Request) (*model.Product, error) {
//	req, err := getReq[vo.ProductReq](r)
//	if err != nil {
//		return nil, err
//	}
//	if req.Action == "" || req.Product == nil {
//		return nil, fmt.Errorf("缺少参数")
//	}
//
//	if req.Action == "inc" {
//		err = service.ProductServiceImpl.UpsertProduct(req)
//		if err != nil {
//			return nil, err
//		}
//	} else if req.Action == "clear" {
//		err = service.ProductServiceImpl.DeleteProduct(int64(req.Product.Id))
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		err = fmt.Errorf("参数 action : %s 错误", req.Action)
//	}
//
//	return req.Product, err
//}
