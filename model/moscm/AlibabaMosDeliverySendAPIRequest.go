package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosDeliverySendAPIRequest
发货 API请求
alibaba.mos.delivery.send

订单发货填写快递单 */
type AlibabaMosDeliverySendAPIRequest struct {
	model.Params
	// 发货信息
	_deliveryDto *DeliveryDto
}

// New
