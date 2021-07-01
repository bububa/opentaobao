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
type AlitripBtripInvoiceGetAPIRequest struct {
    model.Params
    // 企业id
    _corpId   string
    // 用户id
    _userId   string
}

// 初始化AlitripBtripInvoiceGetAPIRequest对象
func NewAlitripBtripInvoiceGetRequest() *AlitripBtripInvoiceGetAPIRequest{
    return &AlitripBtripInvoiceGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripInvoiceGetAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.invoice.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripInvoiceGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CorpId Setter
// 企业id
func (r *AlitripBtripInvoiceGetAPIRequest) SetCorpId(_corpId string) error {
    r._corpId = _corpId
    r.Set("corp_id", _corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripInvoiceGetAPIRequest) GetCorpId() string {
    return r._corpId
}
// UserId Setter
// 用户id
func (r *AlitripBtripInvoiceGetAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlitripBtripInvoiceGetAPIRequest) GetUserId() string {
    return r._userId
}
