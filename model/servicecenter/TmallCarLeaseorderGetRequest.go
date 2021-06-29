package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取租赁订单信息 API请求
tmall.car.leaseorder.get

卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家
*/
type TmallCarLeaseorderGetRequest struct {
    model.Params
    // 订单号
    _orderId   int64
}

// 初始化TmallCarLeaseorderGetRequest对象
func NewTmallCarLeaseorderGetRequest() *TmallCarLeaseorderGetRequest{
    return &TmallCarLeaseorderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallCarLeaseorderGetRequest) GetApiMethodName() string {
    return "tmall.car.leaseorder.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallCarLeaseorderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单号
func (r *TmallCarLeaseorderGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TmallCarLeaseorderGetRequest) GetOrderId() int64 {
    return r._orderId
}
