package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机掉货成功通知 API请求
alibaba.retail.device.trade.ship

贩卖机发货
*/
type AlibabaRetailDeviceTradeShipRequest struct {
    model.Params
    // 设备类型
    _deviceType   string
    // 设备ID
    _deviceId   string
    // 内部订单编号
    _tradeNo   string
    // 详情
    _shipDetailList   []ShipDetailDTO
    // 选项
    _orderUpdateOption   *OrderUpdateOption
}

// 初始化AlibabaRetailDeviceTradeShipRequest对象
func NewAlibabaRetailDeviceTradeShipRequest() *AlibabaRetailDeviceTradeShipRequest{
    return &AlibabaRetailDeviceTradeShipRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceTradeShipRequest) GetApiMethodName() string {
    return "alibaba.retail.device.trade.ship"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceTradeShipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceType Setter
// 设备类型
func (r *AlibabaRetailDeviceTradeShipRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetDeviceType() string {
    return r._deviceType
}
// DeviceId Setter
// 设备ID
func (r *AlibabaRetailDeviceTradeShipRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetDeviceId() string {
    return r._deviceId
}
// TradeNo Setter
// 内部订单编号
func (r *AlibabaRetailDeviceTradeShipRequest) SetTradeNo(_tradeNo string) error {
    r._tradeNo = _tradeNo
    r.Set("trade_no", _tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetTradeNo() string {
    return r._tradeNo
}
// ShipDetailList Setter
// 详情
func (r *AlibabaRetailDeviceTradeShipRequest) SetShipDetailList(_shipDetailList []ShipDetailDTO) error {
    r._shipDetailList = _shipDetailList
    r.Set("ship_detail_list", _shipDetailList)
    return nil
}

// ShipDetailList Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetShipDetailList() []ShipDetailDTO {
    return r._shipDetailList
}
// OrderUpdateOption Setter
// 选项
func (r *AlibabaRetailDeviceTradeShipRequest) SetOrderUpdateOption(_orderUpdateOption *OrderUpdateOption) error {
    r._orderUpdateOption = _orderUpdateOption
    r.Set("order_update_option", _orderUpdateOption)
    return nil
}

// OrderUpdateOption Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetOrderUpdateOption() *OrderUpdateOption {
    return r._orderUpdateOption
}
