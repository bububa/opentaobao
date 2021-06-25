package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单查询 APIRequest
taobao.qimen.transferorder.query

调拨单查询
*/
type TaobaoQimenTransferorderQueryRequest struct {
    model.Params

    // 
    request   *TaobaoQimenTransferorderQueryStruct 

}

func NewTaobaoQimenTransferorderQueryRequest() *TaobaoQimenTransferorderQueryRequest{
    return &TaobaoQimenTransferorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenTransferorderQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.transferorder.query"
}

func (r TaobaoQimenTransferorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenTransferorderQueryRequest) SetRequest(request *TaobaoQimenTransferorderQueryStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenTransferorderQueryRequest) GetRequest() *TaobaoQimenTransferorderQueryStruct {
    return r.request
}

