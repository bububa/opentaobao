package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据发票抬头搜索发票 API请求
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

// 初始化AlitripBtripInvoiceSearchRequest对象
func NewAlitripBtripInvoiceSearchRequest() *AlitripBtripInvoiceSearchRequest{
    return &AlitripBtripInvoiceSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceSearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CorpId Setter
// 企业id
func (r *AlitripBtripInvoiceSearchRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripInvoiceSearchRequest) GetCorpId() string {
    return r.corpId
}
// UserId Setter
// 用户id
func (r *AlitripBtripInvoiceSearchRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlitripBtripInvoiceSearchRequest) GetUserId() string {
    return r.userId
}
// Title Setter
// 发票抬头
func (r *AlitripBtripInvoiceSearchRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlitripBtripInvoiceSearchRequest) GetTitle() string {
    return r.title
}
