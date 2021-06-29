package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询状态接口 API请求
alibaba.security.jaq.rp.status

聚安全实人认证查询状态接口
*/
type AlibabaSecurityJaqRpStatusRequest struct {
    model.Params
    // 账号id
    _accountId   string
    // 凭据id
    _ticketId   string
    // 客户端来源
    _source   string
    // 业务来源
    _biz   string
}

// 初始化AlibabaSecurityJaqRpStatusRequest对象
func NewAlibabaSecurityJaqRpStatusRequest() *AlibabaSecurityJaqRpStatusRequest{
    return &AlibabaSecurityJaqRpStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpStatusRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.status"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 账号id
func (r *AlibabaSecurityJaqRpStatusRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetAccountId() string {
    return r._accountId
}
// TicketId Setter
// 凭据id
func (r *AlibabaSecurityJaqRpStatusRequest) SetTicketId(_ticketId string) error {
    r._ticketId = _ticketId
    r.Set("ticket_id", _ticketId)
    return nil
}

// TicketId Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetTicketId() string {
    return r._ticketId
}
// Source Setter
// 客户端来源
func (r *AlibabaSecurityJaqRpStatusRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetSource() string {
    return r._source
}
// Biz Setter
// 业务来源
func (r *AlibabaSecurityJaqRpStatusRequest) SetBiz(_biz string) error {
    r._biz = _biz
    r.Set("biz", _biz)
    return nil
}

// Biz Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetBiz() string {
    return r._biz
}
