package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单创建接口 APIRequest
taobao.qimen.entryorder.create

ERP调用接口，创建入库单;
*/
type TaobaoQimenEntryorderCreateRequest struct {
    model.Params

    // 
    request   *EntryOrderCreateRequest 

}

func NewTaobaoQimenEntryorderCreateRequest() *TaobaoQimenEntryorderCreateRequest{
    return &TaobaoQimenEntryorderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenEntryorderCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.entryorder.create"
}

func (r TaobaoQimenEntryorderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenEntryorderCreateRequest) SetRequest(request *EntryOrderCreateRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenEntryorderCreateRequest) GetRequest() *EntryOrderCreateRequest {
    return r.request
}

