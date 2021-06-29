package consignplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟发货工作台取消包裹以及订单 APIRequest
cainiao.consignplatform.order.cancel

菜鸟发货工作台，商家或者isv通过api取消包裹、回收单号，如果是裹裹运力会取消小件员上门。最后删除订单信息。
*/
type CainiaoConsignplatformOrderCancelRequest struct {
    model.Params

    // 取消参数
    cancelRequest   *OrderCancelRequest 

}

func NewCainiaoConsignplatformOrderCancelRequest() *CainiaoConsignplatformOrderCancelRequest{
    return &CainiaoConsignplatformOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoConsignplatformOrderCancelRequest) GetApiMethodName() string {
    return "cainiao.consignplatform.order.cancel"
}

func (r CainiaoConsignplatformOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoConsignplatformOrderCancelRequest) SetCancelRequest(cancelRequest *OrderCancelRequest) error {
    r.cancelRequest = cancelRequest
    r.Set("cancel_request", cancelRequest)
    return nil
}

func (r CainiaoConsignplatformOrderCancelRequest) GetCancelRequest() *OrderCancelRequest {
    return r.cancelRequest
}

