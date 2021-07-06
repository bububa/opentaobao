package alihealth2

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.modifyrestime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalModifyrestimeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
