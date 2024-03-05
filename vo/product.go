package vo

import "wxcloudrun-golang/db/model"

type ProductReq struct {
	Action  string         `json:"action"`
	Product *model.Product `json:"product"`
}
