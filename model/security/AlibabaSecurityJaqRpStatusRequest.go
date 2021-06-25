package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询状态接口 APIRequest
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

func NewAlibabaSecurityJaqRpStatusRequest() *AlibabaSecurityJaqRpStatusRequest{
    return &AlibabaSecurityJaqRpStatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpStatusRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.status"
}

func (r AlibabaSecurityJaqRpStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpStatusRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaSecurityJaqRpStatusRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaSecurityJaqRpStatusRequest) SetTicketId(ticketId string) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

func (r AlibabaSecurityJaqRpStatusRequest) GetTicketId() string {
    return r.ticketId
}

func (r *AlibabaSecurityJaqRpStatusRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaSecurityJaqRpStatusRequest) GetSource() string {
    return r.source
}

func (r *AlibabaSecurityJaqRpStatusRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

func (r AlibabaSecurityJaqRpStatusRequest) GetBiz() string {
    return r.biz
}

