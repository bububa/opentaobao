package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAppraiseOrderPerformAPIRequest
闲鱼验货担保服务商订单履约V1 API请求
alibaba.idle.appraise.order.perform

闲鱼验货担保业务中,外部服务商作为鉴定方 需要驱动交易节点变化 */
type AlibabaIdleAppraiseOrderPerformAPIRequest struct {
	model.Params
	// AppraiseOrderSynDto
	_appraiseOrderSynDto *AppraiseOrderSynDto
}

// New
