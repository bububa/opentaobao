package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单据取消接口 APIRequest
taobao.qimen.order.cancel

ERP调用奇门的接口,取消创建单据操作。场景介绍：ERP主动发起取消某些创建的单据。如入库单、出库单、退货单等；所有的场景
*/
type TaobaoQimenOrderCancelRequest struct {
    model.Params

    // 
    request   *OrderCancelRequest 

}

func NewTaobaoQimenOrderCancelRequest() *TaobaoQimenOrderCancelRequest{
    return &TaobaoQimenOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenOrderCancelRequest) GetApiMethodName() string {
    return "taobao.qimen.order.cancel"
}

func (r TaobaoQimenOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenOrderCancelRequest) SetRequest(request *OrderCancelRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenOrderCancelRequest) GetRequest() *OrderCancelRequest {
    return r.request
}

