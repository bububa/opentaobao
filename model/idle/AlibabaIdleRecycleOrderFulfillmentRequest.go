package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼回收订单履约V1 APIRequest
alibaba.idle.recycle.order.fulfillment

外部回收商针对自有回收订单的履行
*/
type AlibabaIdleRecycleOrderFulfillmentRequest struct {
    model.Params

    // 订单同步入参
    param0   *RecycleOrderSynDto 

}

func NewAlibabaIdleRecycleOrderFulfillmentRequest() *AlibabaIdleRecycleOrderFulfillmentRequest{
    return &AlibabaIdleRecycleOrderFulfillmentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleRecycleOrderFulfillmentRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.fulfillment"
}

func (r AlibabaIdleRecycleOrderFulfillmentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleRecycleOrderFulfillmentRequest) SetParam0(param0 *RecycleOrderSynDto) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaIdleRecycleOrderFulfillmentRequest) GetParam0() *RecycleOrderSynDto {
    return r.param0
}

