package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest 检测服务-服务单状态流转 API请求
// taobao.auction.zc.vehicle.detect.status.process
//
// 检测服务-服务单状态流转
type TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest struct {
	model.Params
	// 服务入参
	_vehicleServerOrderInfo4Top *VehicleServerOrderInfo4top
}

// NewTaobaoAuctionZcVehicleDetectStatusProcessRequest 初始化TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest对象
func NewTaobaoAuctionZcVehicleDetectStatusProcessRequest() *TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest {
	return &TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) GetApiMethodName() string {
	return "taobao.auction.zc.vehicle.detect.status.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVehicleServerOrderInfo4Top is VehicleServerOrderInfo4Top Setter
// 服务入参
func (r *TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) SetVehicleServerOrderInfo4Top(_vehicleServerOrderInfo4Top *VehicleServerOrderInfo4top) error {
	r._vehicleServerOrderInfo4Top = _vehicleServerOrderInfo4Top
	r.Set("vehicle_server_order_info4_top", _vehicleServerOrderInfo4Top)
	return nil
}

// GetVehicleServerOrderInfo4Top VehicleServerOrderInfo4Top Getter
func (r TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) GetVehicleServerOrderInfo4Top() *VehicleServerOrderInfo4top {
	return r._vehicleServerOrderInfo4Top
}
