package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
o2o查看处方图片 APIRequest
alibaba.alihealth.nr.rx.queryimage

o2o商家查看处方图片，包括电子图片与纸质图片
*/
type AlibabaAlihealthNrRxQueryimageRequest struct {
    model.Params

    // 订单编号
    orderId   int64 

}

func NewAlibabaAlihealthNrRxQueryimageRequest() *AlibabaAlihealthNrRxQueryimageRequest{
    return &AlibabaAlihealthNrRxQueryimageRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthNrRxQueryimageRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.rx.queryimage"
}

func (r AlibabaAlihealthNrRxQueryimageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthNrRxQueryimageRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthNrRxQueryimageRequest) GetOrderId() int64 {
    return r.orderId
}

