package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleAppraiseOrderPerform 闲鱼验货宝服务商订单履约
// alibaba.idle.appraise.order.perform
//
// 闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化
func AlibabaIdleAppraiseOrderPerform(clt *core.SDKClient, req *idle.AlibabaIdleAppraiseOrderPerformAPIRequest, resp *idle.AlibabaIdleAppraiseOrderPerformAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
