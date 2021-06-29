package miniappopen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建小程序投放计划 APIRequest
taobao.miniapp.distribution.order.create

帮助商家，创建小程序的投放计划。
*/
type TaobaoMiniappDistributionOrderCreateRequest struct {
    model.Params

    // 投放计划信息
    orderRequest   *DistributionOrderSaveOpenRequest 

}

func NewTaobaoMiniappDistributionOrderCreateRequest() *TaobaoMiniappDistributionOrderCreateRequest{
    return &TaobaoMiniappDistributionOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoMiniappDistributionOrderCreateRequest) GetApiMethodName() string {
    return "taobao.miniapp.distribution.order.create"
}

func (r TaobaoMiniappDistributionOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoMiniappDistributionOrderCreateRequest) SetOrderRequest(orderRequest *DistributionOrderSaveOpenRequest) error {
    r.orderRequest = orderRequest
    r.Set("order_request", orderRequest)
    return nil
}

func (r TaobaoMiniappDistributionOrderCreateRequest) GetOrderRequest() *DistributionOrderSaveOpenRequest {
    return r.orderRequest
}

