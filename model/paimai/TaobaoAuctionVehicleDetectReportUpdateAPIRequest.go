package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionvehicledetectreportupdateAPIRequest 检测服务-服务单报告信息更新 API请求
// taobao.auction.vehicle.detect.report.update
//
// 检测服务-服务单报告信息更新
type TaobaoauctionvehicledetectreportupdateAPIRequest struct {
	model.Params
	// 服务入参
	_vehicleDetectServerReport4Top *VehicleDetectServerReport4top
}

// NewTaobaoauctionvehicledetectreportupdateRequest 初始化TaobaoauctionvehicledetectreportupdateAPIRequest对象
func NewTaobaoauctionvehicledetectreportupdateRequest() *TaobaoauctionvehicledetectreportupdateAPIRequest {
	return &TaobaoauctionvehicledetectreportupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctionvehicledetectreportupdateAPIRequest) GetApiMethodName() string {
	return "taobao.auction.vehicle.detect.report.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctionvehicledetectreportupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctionvehicledetectreportupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVehicleDetectServerReport4Top is VehicleDetectServerReport4Top Setter
// 服务入参
func (r *TaobaoauctionvehicledetectreportupdateAPIRequest) SetVehicleDetectServerReport4Top(_vehicleDetectServerReport4Top *VehicleDetectServerReport4top) error {
	r._vehicleDetectServerReport4Top = _vehicleDetectServerReport4Top
	r.Set("vehicle_detect_server_report4_top", _vehicleDetectServerReport4Top)
	return nil
}

// GetVehicleDetectServerReport4Top VehicleDetectServerReport4Top Getter
func (r TaobaoauctionvehicledetectreportupdateAPIRequest) GetVehicleDetectServerReport4Top() *VehicleDetectServerReport4top {
	return r._vehicleDetectServerReport4Top
}
