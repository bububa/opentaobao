package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbookingreservetreatAPIRequest 确认就诊 API请求
// alibaba.alihealth.booking.reserve.treat
//
// 统一预约平台，ISV确认就诊
type AlibabaalihealthbookingreservetreatAPIRequest struct {
	model.Params
	// treat
	_treat *IsvReserveRequest
}

// NewAlibabaalihealthbookingreservetreatRequest 初始化AlibabaalihealthbookingreservetreatAPIRequest对象
func NewAlibabaalihealthbookingreservetreatRequest() *AlibabaalihealthbookingreservetreatAPIRequest {
	return &AlibabaalihealthbookingreservetreatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbookingreservetreatAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.treat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbookingreservetreatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbookingreservetreatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTreat is Treat Setter
// treat
func (r *AlibabaalihealthbookingreservetreatAPIRequest) SetTreat(_treat *IsvReserveRequest) error {
	r._treat = _treat
	r.Set("treat", _treat)
	return nil
}

// GetTreat Treat Getter
func (r AlibabaalihealthbookingreservetreatAPIRequest) GetTreat() *IsvReserveRequest {
	return r._treat
}
