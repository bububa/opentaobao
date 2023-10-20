package moscm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosDeliverySendAPIRequest 发货 API请求
// alibaba.mos.delivery.send
//
// 订单发货填写快递单
type AlibabaMosDeliverySendAPIRequest struct {
	model.Params
	// 发货信息
	_deliveryDto *DeliveryDto
}

// NewAlibabaMosDeliverySendRequest 初始化AlibabaMosDeliverySendAPIRequest对象
func NewAlibabaMosDeliverySendRequest() *AlibabaMosDeliverySendAPIRequest {
	return &AlibabaMosDeliverySendAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosDeliverySendAPIRequest) Reset() {
	r._deliveryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosDeliverySendAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.delivery.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosDeliverySendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosDeliverySendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryDto is DeliveryDto Setter
// 发货信息
func (r *AlibabaMosDeliverySendAPIRequest) SetDeliveryDto(_deliveryDto *DeliveryDto) error {
	r._deliveryDto = _deliveryDto
	r.Set("delivery_dto", _deliveryDto)
	return nil
}

// GetDeliveryDto DeliveryDto Getter
func (r AlibabaMosDeliverySendAPIRequest) GetDeliveryDto() *DeliveryDto {
	return r._deliveryDto
}

var poolAlibabaMosDeliverySendAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosDeliverySendRequest()
	},
}

// GetAlibabaMosDeliverySendRequest 从 sync.Pool 获取 AlibabaMosDeliverySendAPIRequest
func GetAlibabaMosDeliverySendAPIRequest() *AlibabaMosDeliverySendAPIRequest {
	return poolAlibabaMosDeliverySendAPIRequest.Get().(*AlibabaMosDeliverySendAPIRequest)
}

// ReleaseAlibabaMosDeliverySendAPIRequest 将 AlibabaMosDeliverySendAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosDeliverySendAPIRequest(v *AlibabaMosDeliverySendAPIRequest) {
	v.Reset()
	poolAlibabaMosDeliverySendAPIRequest.Put(v)
}
