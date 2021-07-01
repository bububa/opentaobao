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

// NewAlibabaMosDeliverySendRequest 初始化AlibabaMosDeliverySendAPIRequest对象
func NewAlibabaMosDeliverySendRequest() *AlibabaMosDeliverySendAPIRequest {
	return &AlibabaMosDeliverySendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosDeliverySendAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.delivery.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosDeliverySendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeliveryDto Setter
// 发货信息
func (r *AlibabaMosDeliverySendAPIRequest) SetDeliveryDto(_deliveryDto *DeliveryDto) error {
	r._deliveryDto = _deliveryDto
	r.Set("delivery_dto", _deliveryDto)
	return nil
}

// Get DeliveryDto Getter
func (r AlibabaMosDeliverySendAPIRequest) GetDeliveryDto() *DeliveryDto {
	return r._deliveryDto
}
