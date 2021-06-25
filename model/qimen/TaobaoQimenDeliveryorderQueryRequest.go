package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单查询接口 APIRequest
taobao.qimen.deliveryorder.query

ERP调用奇门的发货单查询接口，查询发货单详情
*/
type TaobaoQimenDeliveryorderQueryRequest struct {
    model.Params

    // 
    request   *DeliveryOrderQueryRequest 

}

func NewTaobaoQimenDeliveryorderQueryRequest() *TaobaoQimenDeliveryorderQueryRequest{
    return &TaobaoQimenDeliveryorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenDeliveryorderQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.deliveryorder.query"
}

func (r TaobaoQimenDeliveryorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenDeliveryorderQueryRequest) SetRequest(request *DeliveryOrderQueryRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenDeliveryorderQueryRequest) GetRequest() *DeliveryOrderQueryRequest {
    return r.request
}

