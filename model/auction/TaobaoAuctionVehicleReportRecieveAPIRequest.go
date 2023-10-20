package auction

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoauctionvehiclereportrecieveAPIRequest 机动车报告回调数据接收 API请求
// taobao.auction.vehicle.report.recieve
//
// 机动车报告同步接收接口
type TaobaoauctionvehiclereportrecieveAPIRequest struct {
	model.Params
	// 拍品id
	_itemId int64
	// 机动车报告数据
	_vehicleTestReportDto *VehicleTestReportDto
}

// NewTaobaoauctionvehiclereportrecieveRequest 初始化TaobaoauctionvehiclereportrecieveAPIRequest对象
func NewTaobaoauctionvehiclereportrecieveRequest() *TaobaoauctionvehiclereportrecieveAPIRequest {
	return &TaobaoauctionvehiclereportrecieveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoauctionvehiclereportrecieveAPIRequest) GetApiMethodName() string {
	return "taobao.auction.vehicle.report.recieve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoauctionvehiclereportrecieveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoauctionvehiclereportrecieveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 拍品id
func (r *TaobaoauctionvehiclereportrecieveAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoauctionvehiclereportrecieveAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetVehicleTestReportDto is VehicleTestReportDto Setter
// 机动车报告数据
func (r *TaobaoauctionvehiclereportrecieveAPIRequest) SetVehicleTestReportDto(_vehicleTestReportDto *VehicleTestReportDto) error {
	r._vehicleTestReportDto = _vehicleTestReportDto
	r.Set("vehicle_test_report_dto", _vehicleTestReportDto)
	return nil
}

// GetVehicleTestReportDto VehicleTestReportDto Getter
func (r TaobaoauctionvehiclereportrecieveAPIRequest) GetVehicleTestReportDto() *VehicleTestReportDto {
	return r._vehicleTestReportDto
}
