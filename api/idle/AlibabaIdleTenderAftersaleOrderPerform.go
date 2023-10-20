package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderaftersaleorderperform 闲鱼帮卖售后订单履约
// alibaba.idle.tender.aftersale.order.perform
//
// 闲鱼帮卖售后订单履约
func Alibabaidletenderaftersaleorderperform(clt *core.SDKClient, req *idle.AlibabaidletenderaftersaleorderperformAPIRequest, session string) (*idle.AlibabaidletenderaftersaleorderperformAPIResponse, error) {
	var resp idle.AlibabaidletenderaftersaleorderperformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
