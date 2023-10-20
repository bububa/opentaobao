package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCreateDoAPIRequest top创建DO/RT接口 API请求
// alibaba.pur.create.do
//
// 创建发货单,先创建DO，异步创建RT
type AlibabaPurCreateDoAPIRequest struct {
	model.Params
	// 发货单
	_deliveryTopDTO *DeliveryTopDto
}

// NewAlibabaPurCreateDoRequest 初始化AlibabaPurCreateDoAPIRequest对象
func NewAlibabaPurCreateDoRequest() *AlibabaPurCreateDoAPIRequest {
	return &AlibabaPurCreateDoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurCreateDoAPIRequest) Reset() {
	r._deliveryTopDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurCreateDoAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.create.do"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurCreateDoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurCreateDoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryTopDTO is DeliveryTopDTO Setter
// 发货单
func (r *AlibabaPurCreateDoAPIRequest) SetDeliveryTopDTO(_deliveryTopDTO *DeliveryTopDto) error {
	r._deliveryTopDTO = _deliveryTopDTO
	r.Set("delivery_top_d_t_o", _deliveryTopDTO)
	return nil
}

// GetDeliveryTopDTO DeliveryTopDTO Getter
func (r AlibabaPurCreateDoAPIRequest) GetDeliveryTopDTO() *DeliveryTopDto {
	return r._deliveryTopDTO
}

var poolAlibabaPurCreateDoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurCreateDoRequest()
	},
}

// GetAlibabaPurCreateDoRequest 从 sync.Pool 获取 AlibabaPurCreateDoAPIRequest
func GetAlibabaPurCreateDoAPIRequest() *AlibabaPurCreateDoAPIRequest {
	return poolAlibabaPurCreateDoAPIRequest.Get().(*AlibabaPurCreateDoAPIRequest)
}

// ReleaseAlibabaPurCreateDoAPIRequest 将 AlibabaPurCreateDoAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurCreateDoAPIRequest(v *AlibabaPurCreateDoAPIRequest) {
	v.Reset()
	poolAlibabaPurCreateDoAPIRequest.Put(v)
}
