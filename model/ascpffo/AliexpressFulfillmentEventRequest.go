package ascpffo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AE履约事件处理 APIRequest
aliexpress.fulfillment.event

AE用 履约底层声明发货能力
*/
type AliexpressFulfillmentEventRequest struct {
    model.Params

    // 入参对象
    param   *FulfillmentOrderStatusUpdateRequest 

}

func NewAliexpressFulfillmentEventRequest() *AliexpressFulfillmentEventRequest{
    return &AliexpressFulfillmentEventRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressFulfillmentEventRequest) GetApiMethodName() string {
    return "aliexpress.fulfillment.event"
}

func (r AliexpressFulfillmentEventRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressFulfillmentEventRequest) SetParam(param *FulfillmentOrderStatusUpdateRequest) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AliexpressFulfillmentEventRequest) GetParam() *FulfillmentOrderStatusUpdateRequest {
    return r.param
}

