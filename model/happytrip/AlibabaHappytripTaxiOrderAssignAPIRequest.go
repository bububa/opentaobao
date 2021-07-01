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
type AlibabaHappytripTaxiOrderAssignAPIRequest struct {
    model.Params
    // 供应商订单号
    _orderId   string
}

// 初始化AlibabaHappytripTaxiOrderAssignAPIRequest对象
func NewAlibabaHappytripTaxiOrderAssignRequest() *AlibabaHappytripTaxiOrderAssignAPIRequest{
    return &AlibabaHappytripTaxiOrderAssignAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderAssignAPIRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.assign"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderAssignAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiOrderAssignAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderAssignAPIRequest) GetOrderId() string {
    return r._orderId
}
