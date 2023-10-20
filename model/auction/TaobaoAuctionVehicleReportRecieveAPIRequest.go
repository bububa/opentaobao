package auction

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAuctionVehicleReportRecieveAPIRequest 机动车报告回调数据接收 API请求
// taobao.auction.vehicle.report.recieve
//
// 机动车报告同步接收接口
type TaobaoAuctionVehicleReportRecieveAPIRequest struct {
	model.Params
	// 拍品id
	_itemId int64
	// 机动车报告数据
	_vehicleTestReportDto *VehicleTestReportDto
}

// NewTaobaoAuctionVehicleReportRecieveRequest 初始化TaobaoAuctionVehicleReportRecieveAPIRequest对象
func NewTaobaoAuctionVehicleReportRecieveRequest() *TaobaoAuctionVehicleReportRecieveAPIRequest {
	return &TaobaoAuctionVehicleReportRecieveAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAuctionVehicleReportRecieveAPIRequest) Reset() {
	r._itemId = 0
	r._vehicleTestReportDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionVehicleReportRecieveAPIRequest) GetApiMethodName() string {
	return "taobao.auction.vehicle.report.recieve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionVehicleReportRecieveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAuctionVehicleReportRecieveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 拍品id
func (r *TaobaoAuctionVehicleReportRecieveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAuctionVehicleReportRecieveAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetVehicleTestReportDto is VehicleTestReportDto Setter
// 机动车报告数据
func (r *TaobaoAuctionVehicleReportRecieveAPIRequest) SetVehicleTestReportDto(_vehicleTestReportDto *VehicleTestReportDto) error {
	r._vehicleTestReportDto = _vehicleTestReportDto
	r.Set("vehicle_test_report_dto", _vehicleTestReportDto)
	return nil
}

// GetVehicleTestReportDto VehicleTestReportDto Getter
func (r TaobaoAuctionVehicleReportRecieveAPIRequest) GetVehicleTestReportDto() *VehicleTestReportDto {
	return r._vehicleTestReportDto
}

var poolTaobaoAuctionVehicleReportRecieveAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAuctionVehicleReportRecieveRequest()
	},
}

// GetTaobaoAuctionVehicleReportRecieveRequest 从 sync.Pool 获取 TaobaoAuctionVehicleReportRecieveAPIRequest
func GetTaobaoAuctionVehicleReportRecieveAPIRequest() *TaobaoAuctionVehicleReportRecieveAPIRequest {
	return poolTaobaoAuctionVehicleReportRecieveAPIRequest.Get().(*TaobaoAuctionVehicleReportRecieveAPIRequest)
}

// ReleaseTaobaoAuctionVehicleReportRecieveAPIRequest 将 TaobaoAuctionVehicleReportRecieveAPIRequest 放入 sync.Pool
func ReleaseTaobaoAuctionVehicleReportRecieveAPIRequest(v *TaobaoAuctionVehicleReportRecieveAPIRequest) {
	v.Reset()
	poolTaobaoAuctionVehicleReportRecieveAPIRequest.Put(v)
}
