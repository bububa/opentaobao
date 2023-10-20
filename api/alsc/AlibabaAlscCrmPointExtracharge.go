package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpointextracharge 积分补录
// alibaba.alsc.crm.point.extracharge
//
// 积分补录
func Alibabaalsccrmpointextracharge(clt *core.SDKClient, req *alsc.AlibabaalsccrmpointextrachargeAPIRequest, session string) (*alsc.AlibabaalsccrmpointextrachargeAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpointextrachargeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
