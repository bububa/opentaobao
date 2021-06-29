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
type AlibabaHappytripTaxiDriverLocationGetRequest struct {
    model.Params
    // 供应商订单号
    orderId   string
}

// 初始化AlibabaHappytripTaxiDriverLocationGetRequest对象
func NewAlibabaHappytripTaxiDriverLocationGetRequest() *AlibabaHappytripTaxiDriverLocationGetRequest{
    return &AlibabaHappytripTaxiDriverLocationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiDriverLocationGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.driver.location.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiDriverLocationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiDriverLocationGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiDriverLocationGetRequest) GetOrderId() string {
    return r.orderId
}
