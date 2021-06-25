package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ToB仓储发货 APIRequest
taobao.uop.tob.order.create

ToB仓储发货
*/
type TaobaoUopTobOrderCreateRequest struct {
    model.Params

    // ERP出库对象
    deliveryOrder   *DeliveryOrder 

}

func NewTaobaoUopTobOrderCreateRequest() *TaobaoUopTobOrderCreateRequest{
    return &TaobaoUopTobOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoUopTobOrderCreateRequest) GetApiMethodName() string {
    return "taobao.uop.tob.order.create"
}

func (r TaobaoUopTobOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoUopTobOrderCreateRequest) SetDeliveryOrder(deliveryOrder *DeliveryOrder) error {
    r.deliveryOrder = deliveryOrder
    r.Set("delivery_order", deliveryOrder)
    return nil
}

func (r TaobaoUopTobOrderCreateRequest) GetDeliveryOrder() *DeliveryOrder {
    return r.deliveryOrder
}

