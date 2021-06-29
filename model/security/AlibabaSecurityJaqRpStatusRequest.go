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
    accountId   string
    // 凭据id
    ticketId   string
    // 客户端来源
    source   string
    // 业务来源
    biz   string
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
func (r *AlibabaSecurityJaqRpStatusRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetAccountId() string {
    return r.accountId
}
// TicketId Setter
// 凭据id
func (r *AlibabaSecurityJaqRpStatusRequest) SetTicketId(ticketId string) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

// TicketId Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetTicketId() string {
    return r.ticketId
}
// Source Setter
// 客户端来源
func (r *AlibabaSecurityJaqRpStatusRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetSource() string {
    return r.source
}
// Biz Setter
// 业务来源
func (r *AlibabaSecurityJaqRpStatusRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

// Biz Getter
func (r AlibabaSecurityJaqRpStatusRequest) GetBiz() string {
    return r.biz
}
