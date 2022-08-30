package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardExtrachargeCreate 创建工单额外收费项
// tmall.servicecenter.workcard.extracharge.create
//
// 创建额外收费项
func TmallServicecenterWorkcardExtrachargeCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardExtrachargeCreateAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardExtrachargeCreateAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardExtrachargeCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
