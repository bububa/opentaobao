package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机掉货成功通知 APIRequest
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

func NewAlibabaRetailDeviceTradeShipRequest() *AlibabaRetailDeviceTradeShipRequest{
    return &AlibabaRetailDeviceTradeShipRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailDeviceTradeShipRequest) GetApiMethodName() string {
    return "alibaba.retail.device.trade.ship"
}

func (r AlibabaRetailDeviceTradeShipRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailDeviceTradeShipRequest) SetDeviceType(deviceType string) error {
    r.deviceType = deviceType
    r.Set("device_type", deviceType)
    return nil
}

func (r AlibabaRetailDeviceTradeShipRequest) GetDeviceType() string {
    return r.deviceType
}

func (r *AlibabaRetailDeviceTradeShipRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaRetailDeviceTradeShipRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *AlibabaRetailDeviceTradeShipRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r AlibabaRetailDeviceTradeShipRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *AlibabaRetailDeviceTradeShipRequest) SetShipDetailList(shipDetailList []ShipDetailDTO) error {
    r.shipDetailList = shipDetailList
    r.Set("ship_detail_list", shipDetailList)
    return nil
}

func (r AlibabaRetailDeviceTradeShipRequest) GetShipDetailList() []ShipDetailDTO {
    return r.shipDetailList
}

func (r *AlibabaRetailDeviceTradeShipRequest) SetOrderUpdateOption(orderUpdateOption *OrderUpdateOption) error {
    r.orderUpdateOption = orderUpdateOption
    r.Set("order_update_option", orderUpdateOption)
    return nil
}

func (r AlibabaRetailDeviceTradeShipRequest) GetOrderUpdateOption() *OrderUpdateOption {
    return r.orderUpdateOption
}

