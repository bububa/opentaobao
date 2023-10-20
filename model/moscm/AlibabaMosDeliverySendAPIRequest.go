package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosdeliverysendAPIRequest 发货 API请求
// alibaba.mos.delivery.send
//
// 订单发货填写快递单
type AlibabamosdeliverysendAPIRequest struct {
	model.Params
	// 发货信息
	_deliveryDto *DeliveryDto
}

// NewAlibabamosdeliverysendRequest 初始化AlibabamosdeliverysendAPIRequest对象
func NewAlibabamosdeliverysendRequest() *AlibabamosdeliverysendAPIRequest {
	return &AlibabamosdeliverysendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosdeliverysendAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.delivery.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosdeliverysendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosdeliverysendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryDto is DeliveryDto Setter
// 发货信息
func (r *AlibabamosdeliverysendAPIRequest) SetDeliveryDto(_deliveryDto *DeliveryDto) error {
	r._deliveryDto = _deliveryDto
	r.Set("delivery_dto", _deliveryDto)
	return nil
}

// GetDeliveryDto DeliveryDto Getter
func (r AlibabamosdeliverysendAPIRequest) GetDeliveryDto() *DeliveryDto {
	return r._deliveryDto
}
