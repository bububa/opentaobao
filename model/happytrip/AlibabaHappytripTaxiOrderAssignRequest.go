package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单指派 API请求
alibaba.happytrip.taxi.order.assign

通知供应商订单指派成功
*/
type AlibabaHappytripTaxiOrderAssignRequest struct {
    model.Params
    // 供应商订单号
    orderId   string
}

// 初始化AlibabaHappytripTaxiOrderAssignRequest对象
func NewAlibabaHappytripTaxiOrderAssignRequest() *AlibabaHappytripTaxiOrderAssignRequest{
    return &AlibabaHappytripTaxiOrderAssignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderAssignRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.assign"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderAssignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiOrderAssignRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderAssignRequest) GetOrderId() string {
    return r.orderId
}
