package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Alibabaservicecenterspserviceorderupdate 服务供应链服务单更新
// alibaba.servicecenter.spserviceorder.update
//
// 服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
func Alibabaservicecenterspserviceorderupdate(clt *core.SDKClient, req *tmallservice.AlibabaservicecenterspserviceorderupdateAPIRequest, session string) (*tmallservice.AlibabaservicecenterspserviceorderupdateAPIResponse, error) {
	var resp tmallservice.AlibabaservicecenterspserviceorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
