package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderprepayaction 服务商预付款完成接口
// alibaba.idle.tender.pre.pay.action
//
// 服务商预付款完成接口
func Alibabaidletenderprepayaction(clt *core.SDKClient, req *idle.AlibabaidletenderprepayactionAPIRequest, session string) (*idle.AlibabaidletenderprepayactionAPIResponse, error) {
	var resp idle.AlibabaidletenderprepayactionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
