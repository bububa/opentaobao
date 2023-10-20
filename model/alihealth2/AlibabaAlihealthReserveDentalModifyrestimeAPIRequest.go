package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthReserveDentalModifyrestimeAPIRequest 修改预约时间 API请求
// alibaba.alihealth.reserve.dental.modifyrestime
//
// 修改预约时间
type AlibabaAlihealthReserveDentalModifyrestimeAPIRequest struct {
	model.Params
	// 预约时间
	_reserveTime string
	// 预约单ID
	_reserveId int64
}

// NewAlibabaAlihealthReserveDentalModifyrestimeRequest 初始化AlibabaAlihealthReserveDentalModifyrestimeAPIRequest对象
func NewAlibabaAlihealthReserveDentalModifyrestimeRequest() *AlibabaAlihealthReserveDentalModifyrestimeAPIRequest {
	return &AlibabaAlihealthReserveDentalModifyrestimeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) Reset() {
	r._reserveTime = ""
	r._reserveId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.modifyrestime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveTime is ReserveTime Setter
// 预约时间
func (r *AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) SetReserveTime(_reserveTime string) error {
	r._reserveTime = _reserveTime
	r.Set("reserve_time", _reserveTime)
	return nil
}

// GetReserveTime ReserveTime Getter
func (r AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) GetReserveTime() string {
	return r._reserveTime
}

// SetReserveId is ReserveId Setter
// 预约单ID
func (r *AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) SetReserveId(_reserveId int64) error {
	r._reserveId = _reserveId
	r.Set("reserve_id", _reserveId)
	return nil
}

// GetReserveId ReserveId Getter
func (r AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) GetReserveId() int64 {
	return r._reserveId
}

var poolAlibabaAlihealthReserveDentalModifyrestimeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthReserveDentalModifyrestimeRequest()
	},
}

// GetAlibabaAlihealthReserveDentalModifyrestimeRequest 从 sync.Pool 获取 AlibabaAlihealthReserveDentalModifyrestimeAPIRequest
func GetAlibabaAlihealthReserveDentalModifyrestimeAPIRequest() *AlibabaAlihealthReserveDentalModifyrestimeAPIRequest {
	return poolAlibabaAlihealthReserveDentalModifyrestimeAPIRequest.Get().(*AlibabaAlihealthReserveDentalModifyrestimeAPIRequest)
}

// ReleaseAlibabaAlihealthReserveDentalModifyrestimeAPIRequest 将 AlibabaAlihealthReserveDentalModifyrestimeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthReserveDentalModifyrestimeAPIRequest(v *AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthReserveDentalModifyrestimeAPIRequest.Put(v)
}
