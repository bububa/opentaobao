package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveModifyAPIRequest 修改预约 API请求
// alibaba.alihealth.booking.reserve.modify
//
// 消费医疗统一预约平台，取消预约
type AlibabaAlihealthBookingReserveModifyAPIRequest struct {
	model.Params
	// modify
	_modify *IsvReserveRequest
}

// NewAlibabaAlihealthBookingReserveModifyRequest 初始化AlibabaAlihealthBookingReserveModifyAPIRequest对象
func NewAlibabaAlihealthBookingReserveModifyRequest() *AlibabaAlihealthBookingReserveModifyAPIRequest {
	return &AlibabaAlihealthBookingReserveModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetModify is Modify Setter
// modify
func (r *AlibabaAlihealthBookingReserveModifyAPIRequest) SetModify(_modify *IsvReserveRequest) error {
	r._modify = _modify
	r.Set("modify", _modify)
	return nil
}

// GetModify Modify Getter
func (r AlibabaAlihealthBookingReserveModifyAPIRequest) GetModify() *IsvReserveRequest {
	return r._modify
}
