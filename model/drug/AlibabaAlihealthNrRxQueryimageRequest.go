package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
o2o查看处方图片 API请求
alibaba.alihealth.nr.rx.queryimage

o2o商家查看处方图片，包括电子图片与纸质图片
*/
type AlibabaAlihealthNrRxQueryimageRequest struct {
    model.Params
    // 订单编号
    orderId   int64
}

// 初始化AlibabaAlihealthNrRxQueryimageRequest对象
func NewAlibabaAlihealthNrRxQueryimageRequest() *AlibabaAlihealthNrRxQueryimageRequest{
    return &AlibabaAlihealthNrRxQueryimageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrRxQueryimageRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.rx.queryimage"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrRxQueryimageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单编号
func (r *AlibabaAlihealthNrRxQueryimageRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthNrRxQueryimageRequest) GetOrderId() int64 {
    return r.orderId
}
