package alihealth2

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthBookingReserveModifyAPIRequest) Reset() {
	r._modify = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBookingReserveModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthBookingReserveModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthBookingReserveModifyRequest()
	},
}

// GetAlibabaAlihealthBookingReserveModifyRequest 从 sync.Pool 获取 AlibabaAlihealthBookingReserveModifyAPIRequest
func GetAlibabaAlihealthBookingReserveModifyAPIRequest() *AlibabaAlihealthBookingReserveModifyAPIRequest {
	return poolAlibabaAlihealthBookingReserveModifyAPIRequest.Get().(*AlibabaAlihealthBookingReserveModifyAPIRequest)
}

// ReleaseAlibabaAlihealthBookingReserveModifyAPIRequest 将 AlibabaAlihealthBookingReserveModifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveModifyAPIRequest(v *AlibabaAlihealthBookingReserveModifyAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveModifyAPIRequest.Put(v)
}
