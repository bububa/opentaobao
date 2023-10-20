package paimai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionVehicleDetectReportUpdateAPIRequest 检测服务-服务单报告信息更新 API请求
// taobao.auction.vehicle.detect.report.update
//
// 检测服务-服务单报告信息更新
type TaobaoAuctionVehicleDetectReportUpdateAPIRequest struct {
	model.Params
	// 服务入参
	_vehicleDetectServerReport4Top *VehicleDetectServerReport4Top
}

// NewTaobaoAuctionVehicleDetectReportUpdateRequest 初始化TaobaoAuctionVehicleDetectReportUpdateAPIRequest对象
func NewTaobaoAuctionVehicleDetectReportUpdateRequest() *TaobaoAuctionVehicleDetectReportUpdateAPIRequest {
	return &TaobaoAuctionVehicleDetectReportUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAuctionVehicleDetectReportUpdateAPIRequest) Reset() {
	r._vehicleDetectServerReport4Top = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionVehicleDetectReportUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.auction.vehicle.detect.report.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionVehicleDetectReportUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionVehicleDetectReportUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVehicleDetectServerReport4Top is VehicleDetectServerReport4Top Setter
// 服务入参
func (r *TaobaoAuctionVehicleDetectReportUpdateAPIRequest) SetVehicleDetectServerReport4Top(_vehicleDetectServerReport4Top *VehicleDetectServerReport4Top) error {
	r._vehicleDetectServerReport4Top = _vehicleDetectServerReport4Top
	r.Set("vehicle_detect_server_report4_top", _vehicleDetectServerReport4Top)
	return nil
}

// GetVehicleDetectServerReport4Top VehicleDetectServerReport4Top Getter
func (r TaobaoAuctionVehicleDetectReportUpdateAPIRequest) GetVehicleDetectServerReport4Top() *VehicleDetectServerReport4Top {
	return r._vehicleDetectServerReport4Top
}

var poolTaobaoAuctionVehicleDetectReportUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAuctionVehicleDetectReportUpdateRequest()
	},
}

// GetTaobaoAuctionVehicleDetectReportUpdateRequest 从 sync.Pool 获取 TaobaoAuctionVehicleDetectReportUpdateAPIRequest
func GetTaobaoAuctionVehicleDetectReportUpdateAPIRequest() *TaobaoAuctionVehicleDetectReportUpdateAPIRequest {
	return poolTaobaoAuctionVehicleDetectReportUpdateAPIRequest.Get().(*TaobaoAuctionVehicleDetectReportUpdateAPIRequest)
}

// ReleaseTaobaoAuctionVehicleDetectReportUpdateAPIRequest 将 TaobaoAuctionVehicleDetectReportUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoAuctionVehicleDetectReportUpdateAPIRequest(v *TaobaoAuctionVehicleDetectReportUpdateAPIRequest) {
	v.Reset()
	poolTaobaoAuctionVehicleDetectReportUpdateAPIRequest.Put(v)
}
