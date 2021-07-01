package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleConsignmentOrderPerformAPIRequest
帮卖订单履约 API请求
alibaba.idle.consignment.order.perform

帮卖订单履约，回收商同步订单信息，驱动交易流转 */
type AlibabaIdleConsignmentOrderPerformAPIRequest struct {
	model.Params
	// 帮卖订单同步DTO
	_param *ConsignmentOrderSynDto
}

// New
