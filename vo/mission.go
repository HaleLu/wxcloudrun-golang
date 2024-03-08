package vo

import "wxcloudrun-golang/db/model"

type MissionReq struct {
	Action     string         `json:"action"`
	Mission    *model.Mission `json:"mission"`
	MissionId  uint64         `json:"mission_id"`
	FinishTime int64          `json:"finish_time"`
}
