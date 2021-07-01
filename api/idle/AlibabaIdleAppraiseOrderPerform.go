package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

/* AlibabaIdleAppraiseOrderPerform
闲鱼验货担保服务商订单履约V1
alibaba.idle.appraise.order.perform

闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化 */
func AlibabaIdleAppraiseOrderPerform(clt *core.SDKClient, req *idle.AlibabaIdleAppraiseOrderPerformAPIRequest, session string) (*idle.AlibabaIdleAppraiseOrderPerformAPIResponse, error) {
	var resp idle.AlibabaIdleAppraiseOrderPerformAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
