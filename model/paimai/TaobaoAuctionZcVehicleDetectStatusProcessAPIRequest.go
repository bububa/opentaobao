package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionzcvehicledetectstatusprocessAPIRequest 检测服务-服务单状态流转 API请求
// taobao.auction.zc.vehicle.detect.status.process
//
// 检测服务-服务单状态流转
type TaobaoauctionzcvehicledetectstatusprocessAPIRequest struct {
	model.Params
	// 服务入参
	_vehicleServerOrderInfo4Top *VehicleServerOrderInfo4top
}

// NewTaobaoauctionzcvehicledetectstatusprocessRequest 初始化TaobaoauctionzcvehicledetectstatusprocessAPIRequest对象
func NewTaobaoauctionzcvehicledetectstatusprocessRequest() *TaobaoauctionzcvehicledetectstatusprocessAPIRequest {
	return &TaobaoauctionzcvehicledetectstatusprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctionzcvehicledetectstatusprocessAPIRequest) GetApiMethodName() string {
	return "taobao.auction.zc.vehicle.detect.status.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctionzcvehicledetectstatusprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctionzcvehicledetectstatusprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVehicleServerOrderInfo4Top is VehicleServerOrderInfo4Top Setter
// 服务入参
func (r *TaobaoauctionzcvehicledetectstatusprocessAPIRequest) SetVehicleServerOrderInfo4Top(_vehicleServerOrderInfo4Top *VehicleServerOrderInfo4top) error {
	r._vehicleServerOrderInfo4Top = _vehicleServerOrderInfo4Top
	r.Set("vehicle_server_order_info4_top", _vehicleServerOrderInfo4Top)
	return nil
}

// GetVehicleServerOrderInfo4Top VehicleServerOrderInfo4Top Getter
func (r TaobaoauctionzcvehicledetectstatusprocessAPIRequest) GetVehicleServerOrderInfo4Top() *VehicleServerOrderInfo4top {
	return r._vehicleServerOrderInfo4Top
}
