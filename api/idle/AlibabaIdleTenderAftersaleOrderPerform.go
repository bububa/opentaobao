package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderAftersaleOrderPerform 闲鱼帮卖售后订单履约
// alibaba.idle.tender.aftersale.order.perform
//
// 闲鱼帮卖售后订单履约
func AlibabaIdleTenderAftersaleOrderPerform(clt *core.SDKClient, req *idle.AlibabaIdleTenderAftersaleOrderPerformAPIRequest, session string) (*idle.AlibabaIdleTenderAftersaleOrderPerformAPIResponse, error) {
	var resp idle.AlibabaIdleTenderAftersaleOrderPerformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
