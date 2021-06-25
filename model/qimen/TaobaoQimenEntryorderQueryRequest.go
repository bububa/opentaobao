package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单查询接口 APIRequest
taobao.qimen.entryorder.query

ERP调用接口，查询入库单信息;
*/
type TaobaoQimenEntryorderQueryRequest struct {
    model.Params

    // 
    request   *EntryOrderQueryRequest 

}

func NewTaobaoQimenEntryorderQueryRequest() *TaobaoQimenEntryorderQueryRequest{
    return &TaobaoQimenEntryorderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenEntryorderQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.entryorder.query"
}

func (r TaobaoQimenEntryorderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenEntryorderQueryRequest) SetRequest(request *EntryOrderQueryRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenEntryorderQueryRequest) GetRequest() *EntryOrderQueryRequest {
    return r.request
}

