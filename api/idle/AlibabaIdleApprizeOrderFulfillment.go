package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleApprizeOrderFulfillment 鉴定担保资金订单履约
// alibaba.idle.apprize.order.fulfillment
//
// 服务商针对自己的服务订单进行履约
func AlibabaIdleApprizeOrderFulfillment(clt *core.SDKClient, req *idle.AlibabaIdleApprizeOrderFulfillmentAPIRequest, resp *idle.AlibabaIdleApprizeOrderFulfillmentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
