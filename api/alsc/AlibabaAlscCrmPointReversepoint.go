package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpointreversepoint 积分消费回退
// alibaba.alsc.crm.point.reversepoint
//
// 积分消费回退
func Alibabaalsccrmpointreversepoint(clt *core.SDKClient, req *alsc.AlibabaalsccrmpointreversepointAPIRequest, session string) (*alsc.AlibabaalsccrmpointreversepointAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpointreversepointAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
