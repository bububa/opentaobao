package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼无需物流发货 APIRequest
alibaba.idle.order.dummy.send

适用于电子卡券等虚拟商品不需要物流的商品发货。
*/
type AlibabaIdleOrderDummySendRequest struct {
    model.Params

    // 请求
    paramOrderDummySendRequest   *OrderDummySendRequest 

}

func NewAlibabaIdleOrderDummySendRequest() *AlibabaIdleOrderDummySendRequest{
    return &AlibabaIdleOrderDummySendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleOrderDummySendRequest) GetApiMethodName() string {
    return "alibaba.idle.order.dummy.send"
}

func (r AlibabaIdleOrderDummySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleOrderDummySendRequest) SetParamOrderDummySendRequest(paramOrderDummySendRequest *OrderDummySendRequest) error {
    r.paramOrderDummySendRequest = paramOrderDummySendRequest
    r.Set("param_order_dummy_send_request", paramOrderDummySendRequest)
    return nil
}

func (r AlibabaIdleOrderDummySendRequest) GetParamOrderDummySendRequest() *OrderDummySendRequest {
    return r.paramOrderDummySendRequest
}

