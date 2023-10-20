package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleConsignmentOrderGet 闲鱼帮卖订单查询
// alibaba.idle.consignment.order.get
//
// 闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
func AlibabaIdleConsignmentOrderGet(clt *core.SDKClient, req *idle.AlibabaIdleConsignmentOrderGetAPIRequest, resp *idle.AlibabaIdleConsignmentOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
