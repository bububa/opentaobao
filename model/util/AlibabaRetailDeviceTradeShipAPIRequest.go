package util

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceTradeShipAPIRequest 贩卖机掉货成功通知 API请求
// alibaba.retail.device.trade.ship
//
// 贩卖机发货
type AlibabaRetailDeviceTradeShipAPIRequest struct {
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

// NewAlibabaRetailDeviceTradeShipRequest 初始化AlibabaRetailDeviceTradeShipAPIRequest对象
func NewAlibabaRetailDeviceTradeShipRequest() *AlibabaRetailDeviceTradeShipAPIRequest {
	return &AlibabaRetailDeviceTradeShipAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailDeviceTradeShipAPIRequest) Reset() {
	r._shipDetailList = r._shipDetailList[:0]
	r._deviceType = ""
	r._deviceId = ""
	r._tradeNo = ""
	r._orderUpdateOption = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.trade.ship"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShipDetailList is ShipDetailList Setter
// 详情
func (r *AlibabaRetailDeviceTradeShipAPIRequest) SetShipDetailList(_shipDetailList []ShipDetailDto) error {
	r._shipDetailList = _shipDetailList
	r.Set("ship_detail_list", _shipDetailList)
	return nil
}

// GetShipDetailList ShipDetailList Getter
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetShipDetailList() []ShipDetailDto {
	return r._shipDetailList
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *AlibabaRetailDeviceTradeShipAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *AlibabaRetailDeviceTradeShipAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetTradeNo is TradeNo Setter
// 内部订单编号
func (r *AlibabaRetailDeviceTradeShipAPIRequest) SetTradeNo(_tradeNo string) error {
	r._tradeNo = _tradeNo
	r.Set("trade_no", _tradeNo)
	return nil
}

// GetTradeNo TradeNo Getter
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetTradeNo() string {
	return r._tradeNo
}

// SetOrderUpdateOption is OrderUpdateOption Setter
// 选项
func (r *AlibabaRetailDeviceTradeShipAPIRequest) SetOrderUpdateOption(_orderUpdateOption *OrderUpdateOption) error {
	r._orderUpdateOption = _orderUpdateOption
	r.Set("order_update_option", _orderUpdateOption)
	return nil
}

// GetOrderUpdateOption OrderUpdateOption Getter
func (r AlibabaRetailDeviceTradeShipAPIRequest) GetOrderUpdateOption() *OrderUpdateOption {
	return r._orderUpdateOption
}

var poolAlibabaRetailDeviceTradeShipAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailDeviceTradeShipRequest()
	},
}

// GetAlibabaRetailDeviceTradeShipRequest 从 sync.Pool 获取 AlibabaRetailDeviceTradeShipAPIRequest
func GetAlibabaRetailDeviceTradeShipAPIRequest() *AlibabaRetailDeviceTradeShipAPIRequest {
	return poolAlibabaRetailDeviceTradeShipAPIRequest.Get().(*AlibabaRetailDeviceTradeShipAPIRequest)
}

// ReleaseAlibabaRetailDeviceTradeShipAPIRequest 将 AlibabaRetailDeviceTradeShipAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailDeviceTradeShipAPIRequest(v *AlibabaRetailDeviceTradeShipAPIRequest) {
	v.Reset()
	poolAlibabaRetailDeviceTradeShipAPIRequest.Put(v)
}
