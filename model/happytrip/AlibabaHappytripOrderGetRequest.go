package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取欢行统一订单模型 APIRequest
alibaba.happytrip.order.get

通过订单id获取欢行统一订单模型数据
*/
type AlibabaHappytripOrderGetRequest struct {
    model.Params

    // 订单id
    orderId   int64 

}

func NewAlibabaHappytripOrderGetRequest() *AlibabaHappytripOrderGetRequest{
    return &AlibabaHappytripOrderGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripOrderGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.order.get"
}

func (r AlibabaHappytripOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripOrderGetRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripOrderGetRequest) GetOrderId() int64 {
    return r.orderId
}

