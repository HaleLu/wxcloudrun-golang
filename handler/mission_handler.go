package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/service"
	"wxcloudrun-golang/vo"
)

// MissionsHandler 商品列表接口
func MissionsHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	fmt.Printf("req:%+v\n", r)
	if r.Method == http.MethodGet {
		missions, err := service.MissionServiceImpl.List()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = missions
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

// MissionHandler 商品接口
func MissionHandler(w http.ResponseWriter, r *http.Request) {
	openId := GetOpenId(r)
	fmt.Println("open_id: " + openId)

	res := &JsonResult{}
	fmt.Printf("req:%+v\n", r)
	if r.Method == http.MethodGet {
		id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		mission, err := service.MissionServiceImpl.Get(uint64(id))
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = mission
		}
	} else if r.Method == http.MethodPost {
		mission, err := modifyMission(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = mission
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

// modifyMission 更新任务
func modifyMission(r *http.Request) (*model.Mission, error) {
	req, err := getReq[vo.MissionReq](r)
	if err != nil {
		return nil, err
	}
	if req.Action == "" || req.Mission == nil {
		return nil, fmt.Errorf("缺少参数")
	}

	if req.Action == "inc" {
		err = service.MissionServiceImpl.Upsert(req.Mission)
		if err != nil {
			return nil, err
		}
	} else if req.Action == "clear" {
		err = service.MissionServiceImpl.Delete(req.Mission.Id)
		if err != nil {
			return nil, err
		}
	} else if req.Action == "finish" {
		err = service.MissionServiceImpl.Finish(GetOpenId(r), req.MissionId, time.Unix(req.FinishTime, 0))
		if err != nil {
			return nil, err
		}
	} else {
		err = fmt.Errorf("参数 action : %s 错误", req.Action)
	}
	return req.Mission, err
}
