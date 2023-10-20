package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderAftersaleOrderPerform 闲鱼帮卖售后订单履约
// alibaba.idle.tender.aftersale.order.perform
//
// 闲鱼帮卖售后订单履约
func AlibabaIdleTenderAftersaleOrderPerform(clt *core.SDKClient, req *idle.AlibabaIdleTenderAftersaleOrderPerformAPIRequest, resp *idle.AlibabaIdleTenderAftersaleOrderPerformAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
