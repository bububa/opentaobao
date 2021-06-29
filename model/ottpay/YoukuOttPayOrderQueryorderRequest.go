package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单 API请求
youku.ott.pay.order.queryorder

通过订单号查询订单信息
*/
type YoukuOttPayOrderQueryorderRequest struct {
    model.Params
    // 订单号
    orderNo   string
}

// 初始化YoukuOttPayOrderQueryorderRequest对象
func NewYoukuOttPayOrderQueryorderRequest() *YoukuOttPayOrderQueryorderRequest{
    return &YoukuOttPayOrderQueryorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryorderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.queryorder"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 订单号
func (r *YoukuOttPayOrderQueryorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r YoukuOttPayOrderQueryorderRequest) GetOrderNo() string {
    return r.orderNo
}
