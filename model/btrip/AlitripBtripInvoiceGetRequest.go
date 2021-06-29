package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户可用发票列表 APIRequest
alitrip.btrip.invoice.get

差旅申请用户获取可用发票列表
*/
type AlitripBtripInvoiceGetRequest struct {
    model.Params

    // 企业id
    corpId   string 

    // 用户id
    userId   string 

}

func NewAlitripBtripInvoiceGetRequest() *AlitripBtripInvoiceGetRequest{
    return &AlitripBtripInvoiceGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripInvoiceGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.get"
}

func (r AlitripBtripInvoiceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripInvoiceGetRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

func (r AlitripBtripInvoiceGetRequest) GetCorpId() string {
    return r.corpId
}

func (r *AlitripBtripInvoiceGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlitripBtripInvoiceGetRequest) GetUserId() string {
    return r.userId
}

