package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleAppraiseOrderQuery 闲鱼验货宝订单详情查询
// alibaba.idle.appraise.order.query
//
// 鉴定商调用该接口获取订单状态
func AlibabaIdleAppraiseOrderQuery(clt *core.SDKClient, req *idle.AlibabaIdleAppraiseOrderQueryAPIRequest, resp *idle.AlibabaIdleAppraiseOrderQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
