package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpointextraconsume 积分补扣
// alibaba.alsc.crm.point.extra.consume
//
// 积分补扣
func Alibabaalsccrmpointextraconsume(clt *core.SDKClient, req *alsc.AlibabaalsccrmpointextraconsumeAPIRequest, session string) (*alsc.AlibabaalsccrmpointextraconsumeAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpointextraconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
