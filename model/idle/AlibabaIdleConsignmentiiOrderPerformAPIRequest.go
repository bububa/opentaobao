package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleConsignmentiiOrderPerformAPIRequest
寄卖V2订单履约 API请求
alibaba.idle.consignmentii.order.perform

寄卖V2订单履约，服务商同步订单信息，驱动交易流转 */
type AlibabaIdleConsignmentiiOrderPerformAPIRequest struct {
	model.Params
	// 同步参数
	_consignmentV2OrderSynDto *ConsignmentV2OrderSynDto
}

// New
