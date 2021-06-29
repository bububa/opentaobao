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
    deviceType   string
    // 设备ID
    deviceId   string
    // 内部订单编号
    tradeNo   string
    // 详情
    shipDetailList   []ShipDetailDTO
    // 选项
    orderUpdateOption   *OrderUpdateOption
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
func (r *AlibabaRetailDeviceTradeShipRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

// DeviceType Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetDeviceType() string {
    return r.deviceType
}
// DeviceId Setter
// 设备ID
func (r *AlibabaRetailDeviceTradeShipRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetDeviceId() string {
    return r.deviceId
}
// TradeNo Setter
// 内部订单编号
func (r *AlibabaRetailDeviceTradeShipRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

// TradeNo Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetTradeNo() string {
    return r.tradeNo
}
// ShipDetailList Setter
// 详情
func (r *AlibabaRetailDeviceTradeShipRequest) SetShipDetailList(shipDetailList []ShipDetailDTO) error {
    r.shipDetailList = shipDetailList
    r.Set("ship_detail_list", shipDetailList)
    return nil
}

// ShipDetailList Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetShipDetailList() []ShipDetailDTO {
    return r.shipDetailList
}
// OrderUpdateOption Setter
// 选项
func (r *AlibabaRetailDeviceTradeShipRequest) SetOrderUpdateOption(orderUpdateOption *OrderUpdateOption) error {
    r.orderUpdateOption = orderUpdateOption
    r.Set("order_update_option", orderUpdateOption)
    return nil
}

// OrderUpdateOption Getter
func (r AlibabaRetailDeviceTradeShipRequest) GetOrderUpdateOption() *OrderUpdateOption {
    return r.orderUpdateOption
}
