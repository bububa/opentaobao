package tmallcar

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
isv查询服务工单详情 APIRequest
tmall.aliauto.service.receipt.get

isv查询服务工单详情
*/
type TmallAliautoServiceReceiptGetRequest struct {
    model.Params

    // 工单号
    receiptId   int64 

}

func NewTmallAliautoServiceReceiptGetRequest() *TmallAliautoServiceReceiptGetRequest{
    return &TmallAliautoServiceReceiptGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallAliautoServiceReceiptGetRequest) GetApiMethodName() string {
    return "tmall.aliauto.service.receipt.get"
}

func (r TmallAliautoServiceReceiptGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallAliautoServiceReceiptGetRequest) SetReceiptId(receiptId int64) error {
    r.receiptId = receiptId
    r.Set("receipt_id", receiptId)
    return nil
}

func (r TmallAliautoServiceReceiptGetRequest) GetReceiptId() int64 {
    return r.receiptId
}

