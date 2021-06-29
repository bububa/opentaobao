package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据发票抬头搜索发票 APIRequest
alitrip.btrip.invoice.search

用户根据发票抬头搜索发票信息
*/
type AlitripBtripInvoiceSearchRequest struct {
    model.Params

    // 企业id
    corpId   string 

    // 用户id
    userId   string 

    // 发票抬头
    title   string 

}

func NewAlitripBtripInvoiceSearchRequest() *AlitripBtripInvoiceSearchRequest{
    return &AlitripBtripInvoiceSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripInvoiceSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.search"
}

func (r AlitripBtripInvoiceSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripInvoiceSearchRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

func (r AlitripBtripInvoiceSearchRequest) GetCorpId() string {
    return r.corpId
}

func (r *AlitripBtripInvoiceSearchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlitripBtripInvoiceSearchRequest) GetUserId() string {
    return r.userId
}

func (r *AlitripBtripInvoiceSearchRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r AlitripBtripInvoiceSearchRequest) GetTitle() string {
    return r.title
}

