package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretaildevicetradeshipAPIRequest 贩卖机掉货成功通知 API请求
// alibaba.retail.device.trade.ship
//
// 贩卖机发货
type AlibabaretaildevicetradeshipAPIRequest struct {
	model.Params
	// 详情
	_shipDetailList []ShipDetailDto
	// 设备类型
	_deviceType string
	// 设备ID
	_deviceId string
	// 内部订单编号
	_tradeNo string
	// 选项
	_orderUpdateOption *OrderUpdateOption
}

// NewAlibabaretaildevicetradeshipRequest 初始化AlibabaretaildevicetradeshipAPIRequest对象
func NewAlibabaretaildevicetradeshipRequest() *AlibabaretaildevicetradeshipAPIRequest {
	return &AlibabaretaildevicetradeshipAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretaildevicetradeshipAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.trade.ship"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretaildevicetradeshipAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretaildevicetradeshipAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShipDetailList is ShipDetailList Setter
// 详情
func (r *AlibabaretaildevicetradeshipAPIRequest) SetShipDetailList(_shipDetailList []ShipDetailDto) error {
	r._shipDetailList = _shipDetailList
	r.Set("ship_detail_list", _shipDetailList)
	return nil
}

// GetShipDetailList ShipDetailList Getter
func (r AlibabaretaildevicetradeshipAPIRequest) GetShipDetailList() []ShipDetailDto {
	return r._shipDetailList
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *AlibabaretaildevicetradeshipAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaretaildevicetradeshipAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaretaildevicetradeshipAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaretaildevicetradeshipAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetTradeNo is TradeNo Setter
// 内部订单编号
func (r *AlibabaretaildevicetradeshipAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabaretaildevicetradeshipAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOrderUpdateOption is OrderUpdateOption Setter
// 选项
func (r *AlibabaretaildevicetradeshipAPIRequest) SetOrderUpdateOption(_orderUpdateOption *OrderUpdateOption) error {
	r._orderUpdateOption = _orderUpdateOption
	r.Set("order_update_option", _orderUpdateOption)
	return nil
}

// GetOrderUpdateOption OrderUpdateOption Getter
func (r AlibabaretaildevicetradeshipAPIRequest) GetOrderUpdateOption() *OrderUpdateOption {
	return r._orderUpdateOption
}
