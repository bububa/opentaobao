package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
司机位置 API请求
alibaba.happytrip.taxi.driver.location.get

获取司机实时位置
*/
type AlibabaHappytripTaxiDriverLocationGetAPIRequest struct {
    model.Params
    // 供应商订单号
    _orderId   string
}

// 初始化AlibabaHappytripTaxiDriverLocationGetAPIRequest对象
func NewAlibabaHappytripTaxiDriverLocationGetRequest() *AlibabaHappytripTaxiDriverLocationGetAPIRequest{
    return &AlibabaHappytripTaxiDriverLocationGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverLocationGetAPIRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.driver.location.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverLocationGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiDriverLocationGetAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiDriverLocationGetAPIRequest) GetOrderId() string {
    return r._orderId
}
