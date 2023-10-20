package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpointconsumepoint 积分抵现
// alibaba.alsc.crm.point.consumepoint
//
// 积分抵现
func Alibabaalsccrmpointconsumepoint(clt *core.SDKClient, req *alsc.AlibabaalsccrmpointconsumepointAPIRequest, session string) (*alsc.AlibabaalsccrmpointconsumepointAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpointconsumepointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
