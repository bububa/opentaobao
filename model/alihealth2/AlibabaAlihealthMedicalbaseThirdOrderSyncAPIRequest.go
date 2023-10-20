package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbasethirdordersyncAPIRequest 第三方订单同步 API请求
// alibaba.alihealth.medicalbase.third.order.sync
//
// 第三方订单同步
type AlibabaalihealthmedicalbasethirdordersyncAPIRequest struct {
	model.Params
	// 请求参数
	_orderRequest *MedicalBaseTopRequestDto
}

// NewAlibabaalihealthmedicalbasethirdordersyncRequest 初始化AlibabaalihealthmedicalbasethirdordersyncAPIRequest对象
func NewAlibabaalihealthmedicalbasethirdordersyncRequest() *AlibabaalihealthmedicalbasethirdordersyncAPIRequest {
	return &AlibabaalihealthmedicalbasethirdordersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthmedicalbasethirdordersyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.third.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthmedicalbasethirdordersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthmedicalbasethirdordersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderRequest is OrderRequest Setter
// 请求参数
func (r *AlibabaalihealthmedicalbasethirdordersyncAPIRequest) SetOrderRequest(_orderRequest *MedicalBaseTopRequestDto) error {
	r._orderRequest = _orderRequest
	r.Set("order_request", _orderRequest)
	return nil
}

// GetOrderRequest OrderRequest Getter
func (r AlibabaalihealthmedicalbasethirdordersyncAPIRequest) GetOrderRequest() *MedicalBaseTopRequestDto {
	return r._orderRequest
}
