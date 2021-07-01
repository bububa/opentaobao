package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderDealrefundAPIRequest
闲鱼无忧购入仓模式服务商退款处理接口 API请求
alibaba.idle.isv.order.dealrefund

闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款 */
type AlibabaIdleIsvOrderDealrefundAPIRequest struct {
	model.Params
	// 退款参数
	_paramAppraiseIsvRefundRequest *AppraiseIsvRefundRequest
}

// New
