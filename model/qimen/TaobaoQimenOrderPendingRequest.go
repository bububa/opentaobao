package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据挂起（恢复）接口 APIRequest
taobao.qimen.order.pending

ERP调用奇门的接口,挂起某些创建的单据;场景介绍：ERP主动发起挂起（恢复）某些创建的单据，如入库单、出库单、退货单等
*/
type TaobaoQimenOrderPendingRequest struct {
    model.Params

    // 
    request   *OrderPendingRequest 

}

func NewTaobaoQimenOrderPendingRequest() *TaobaoQimenOrderPendingRequest{
    return &TaobaoQimenOrderPendingRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenOrderPendingRequest) GetApiMethodName() string {
    return "taobao.qimen.order.pending"
}

func (r TaobaoQimenOrderPendingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenOrderPendingRequest) SetRequest(request *OrderPendingRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenOrderPendingRequest) GetRequest() *OrderPendingRequest {
    return r.request
}

