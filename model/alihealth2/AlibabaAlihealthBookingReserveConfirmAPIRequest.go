package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbookingreserveconfirmAPIRequest 确认预约 API请求
// alibaba.alihealth.booking.reserve.confirm
//
// 确认预约
type AlibabaalihealthbookingreserveconfirmAPIRequest struct {
	model.Params
	// 参数
	_confirm *IsvReserveRequest
}

// NewAlibabaalihealthbookingreserveconfirmRequest 初始化AlibabaalihealthbookingreserveconfirmAPIRequest对象
func NewAlibabaalihealthbookingreserveconfirmRequest() *AlibabaalihealthbookingreserveconfirmAPIRequest {
	return &AlibabaalihealthbookingreserveconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbookingreserveconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbookingreserveconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbookingreserveconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirm is Confirm Setter
// 参数
func (r *AlibabaalihealthbookingreserveconfirmAPIRequest) SetConfirm(_confirm *IsvReserveRequest) error {
	r._confirm = _confirm
	r.Set("confirm", _confirm)
	return nil
}

// GetConfirm Confirm Getter
func (r AlibabaalihealthbookingreserveconfirmAPIRequest) GetConfirm() *IsvReserveRequest {
	return r._confirm
}
