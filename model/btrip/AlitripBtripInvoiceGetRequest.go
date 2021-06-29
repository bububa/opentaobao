package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户可用发票列表 API请求
alitrip.btrip.invoice.get

差旅申请用户获取可用发票列表
*/
type AlitripBtripInvoiceGetRequest struct {
    model.Params
    // 企业id
    _corpId   string
    // 用户id
    _userId   string
}

// 初始化AlitripBtripInvoiceGetRequest对象
func NewAlitripBtripInvoiceGetRequest() *AlitripBtripInvoiceGetRequest{
    return &AlitripBtripInvoiceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CorpId Setter
// 企业id
func (r *AlitripBtripInvoiceGetRequest) SetCorpId(_corpId string) error {
    r._corpId = _corpId
    r.Set("corp_id", _corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripInvoiceGetRequest) GetCorpId() string {
    return r._corpId
}
// UserId Setter
// 用户id
func (r *AlitripBtripInvoiceGetRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlitripBtripInvoiceGetRequest) GetUserId() string {
    return r._userId
}
