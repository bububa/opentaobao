package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单详情查询接口 API请求
alibaba.health.nr.cep.order.query

订单详情查询接口
*/
type AlibabaHealthNrCepOrderQueryRequest struct {
    model.Params
    // 订单号
    orderId   int64
}

// 初始化AlibabaHealthNrCepOrderQueryRequest对象
func NewAlibabaHealthNrCepOrderQueryRequest() *AlibabaHealthNrCepOrderQueryRequest{
    return &AlibabaHealthNrCepOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthNrCepOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.health.nr.cep.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthNrCepOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单号
func (r *AlibabaHealthNrCepOrderQueryRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHealthNrCepOrderQueryRequest) GetOrderId() int64 {
    return r.orderId
}
