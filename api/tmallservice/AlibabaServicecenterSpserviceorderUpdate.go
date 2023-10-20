package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterSpserviceorderUpdate 服务供应链服务单更新
// alibaba.servicecenter.spserviceorder.update
//
// 服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中
func AlibabaServicecenterSpserviceorderUpdate(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterSpserviceorderUpdateAPIRequest, session string) (*tmallservice.AlibabaServicecenterSpserviceorderUpdateAPIResponse, error) {
	var resp tmallservice.AlibabaServicecenterSpserviceorderUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
