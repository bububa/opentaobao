package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单创建 APIRequest
taobao.qimen.transferorder.create

调拨单创建
*/
type TaobaoQimenTransferorderCreateRequest struct {
    model.Params

    // 
    request   *TaobaoQimenTransferorderCreateStruct 

}

func NewTaobaoQimenTransferorderCreateRequest() *TaobaoQimenTransferorderCreateRequest{
    return &TaobaoQimenTransferorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenTransferorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.transferorder.create"
}

func (r TaobaoQimenTransferorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenTransferorderCreateRequest) SetRequest(request *TaobaoQimenTransferorderCreateStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenTransferorderCreateRequest) GetRequest() *TaobaoQimenTransferorderCreateStruct {
    return r.request
}

