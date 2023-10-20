package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthreservedentalmodifyrestimeAPIRequest 修改预约时间 API请求
// alibaba.alihealth.reserve.dental.modifyrestime
//
// 修改预约时间
type AlibabaalihealthreservedentalmodifyrestimeAPIRequest struct {
	model.Params
	// 预约时间
	_reserveTime string
	// 预约单ID
	_reserveId int64
}

// NewAlibabaalihealthreservedentalmodifyrestimeRequest 初始化AlibabaalihealthreservedentalmodifyrestimeAPIRequest对象
func NewAlibabaalihealthreservedentalmodifyrestimeRequest() *AlibabaalihealthreservedentalmodifyrestimeAPIRequest {
	return &AlibabaalihealthreservedentalmodifyrestimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthreservedentalmodifyrestimeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.modifyrestime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthreservedentalmodifyrestimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthreservedentalmodifyrestimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReserveTime is ReserveTime Setter
// 预约时间
func (r *AlibabaalihealthreservedentalmodifyrestimeAPIRequest) SetReserveTime(_reserveTime string) error {
	r._reserveTime = _reserveTime
	r.Set("reserve_time", _reserveTime)
	return nil
}

// GetReserveTime ReserveTime Getter
func (r AlibabaalihealthreservedentalmodifyrestimeAPIRequest) GetReserveTime() string {
	return r._reserveTime
}

// SetReserveId is ReserveId Setter
// 预约单ID
func (r *AlibabaalihealthreservedentalmodifyrestimeAPIRequest) SetReserveId(_reserveId int64) error {
	r._reserveId = _reserveId
	r.Set("reserve_id", _reserveId)
	return nil
}

// GetReserveId ReserveId Getter
func (r AlibabaalihealthreservedentalmodifyrestimeAPIRequest) GetReserveId() int64 {
	return r._reserveId
}
