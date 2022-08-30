package auction

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAuctionVehicleReportRecieveAPIRequest) GetApiMethodName() string {
	return "taobao.auction.vehicle.report.recieve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAuctionVehicleReportRecieveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
