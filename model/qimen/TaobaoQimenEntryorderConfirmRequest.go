package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单确认接口 APIRequest
taobao.qimen.entryorder.confirm

WMS调用接口，回传入库单信息;
*/
type TaobaoQimenEntryorderConfirmRequest struct {
    model.Params

    // 
    request   *EntryOrderConfirmRequest 

}

func NewTaobaoQimenEntryorderConfirmRequest() *TaobaoQimenEntryorderConfirmRequest{
    return &TaobaoQimenEntryorderConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenEntryorderConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.entryorder.confirm"
}

func (r TaobaoQimenEntryorderConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenEntryorderConfirmRequest) SetRequest(request *EntryOrderConfirmRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenEntryorderConfirmRequest) GetRequest() *EntryOrderConfirmRequest {
    return r.request
}

