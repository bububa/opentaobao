package paimai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest 检测服务-服务单状态流转 API请求
// taobao.auction.zc.vehicle.detect.status.process
//
// 检测服务-服务单状态流转
type TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest struct {
	model.Params
	// 服务入参
	_vehicleServerOrderInfo4Top *VehicleServerOrderInfo4Top
}

// NewTaobaoAuctionZcVehicleDetectStatusProcessRequest 初始化TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest对象
func NewTaobaoAuctionZcVehicleDetectStatusProcessRequest() *TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest {
	return &TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) Reset() {
	r._vehicleServerOrderInfo4Top = nil
	r.Params.ToZero()
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
func (r *TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) SetVehicleServerOrderInfo4Top(_vehicleServerOrderInfo4Top *VehicleServerOrderInfo4Top) error {
	r._vehicleServerOrderInfo4Top = _vehicleServerOrderInfo4Top
	r.Set("vehicle_server_order_info4_top", _vehicleServerOrderInfo4Top)
	return nil
}

// GetVehicleServerOrderInfo4Top VehicleServerOrderInfo4Top Getter
func (r TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) GetVehicleServerOrderInfo4Top() *VehicleServerOrderInfo4Top {
	return r._vehicleServerOrderInfo4Top
}

var poolTaobaoAuctionZcVehicleDetectStatusProcessAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAuctionZcVehicleDetectStatusProcessRequest()
	},
}

// GetTaobaoAuctionZcVehicleDetectStatusProcessRequest 从 sync.Pool 获取 TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest
func GetTaobaoAuctionZcVehicleDetectStatusProcessAPIRequest() *TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest {
	return poolTaobaoAuctionZcVehicleDetectStatusProcessAPIRequest.Get().(*TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest)
}

// ReleaseTaobaoAuctionZcVehicleDetectStatusProcessAPIRequest 将 TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest 放入 sync.Pool
func ReleaseTaobaoAuctionZcVehicleDetectStatusProcessAPIRequest(v *TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest) {
	v.Reset()
	poolTaobaoAuctionZcVehicleDetectStatusProcessAPIRequest.Put(v)
}
