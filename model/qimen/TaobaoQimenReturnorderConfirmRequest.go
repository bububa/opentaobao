package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单确认接口 APIRequest
taobao.qimen.returnorder.confirm

taobao.qimen.returnorder.confirm
*/
type TaobaoQimenReturnorderConfirmRequest struct {
    model.Params

    // 
    request   *ReturnOrderConfirmRequest 

}

func NewTaobaoQimenReturnorderConfirmRequest() *TaobaoQimenReturnorderConfirmRequest{
    return &TaobaoQimenReturnorderConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenReturnorderConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.returnorder.confirm"
}

func (r TaobaoQimenReturnorderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenReturnorderConfirmRequest) SetRequest(request *ReturnOrderConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenReturnorderConfirmRequest) GetRequest() *ReturnOrderConfirmRequest {
    return r.request
}

