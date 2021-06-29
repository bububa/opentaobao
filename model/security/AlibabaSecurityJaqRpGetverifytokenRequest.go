package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取认证会话token API请求
alibaba.security.jaq.rp.getverifytoken

聚安全实人认证获取认证会话token
*/
type AlibabaSecurityJaqRpGetverifytokenRequest struct {
    model.Params
    // 账号，强烈建议填写，区别用户的唯一标识
    accountId   string
    // 选填，作为一次验证的唯一标识，每次验证需更换。如果不是验证类型可不填
    ticketId   string
    // 客户端来源
    source   string
    // 业务点
    biz   string
    // 额外信息
    extraData   string
}

// 初始化AlibabaSecurityJaqRpGetverifytokenRequest对象
func NewAlibabaSecurityJaqRpGetverifytokenRequest() *AlibabaSecurityJaqRpGetverifytokenRequest{
    return &AlibabaSecurityJaqRpGetverifytokenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.getverifytoken"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountId Setter
// 账号，强烈建议填写，区别用户的唯一标识
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetAccountId() string {
    return r.accountId
}
// TicketId Setter
// 选填，作为一次验证的唯一标识，每次验证需更换。如果不是验证类型可不填
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetTicketId(ticketId string) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

// TicketId Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetTicketId() string {
    return r.ticketId
}
// Source Setter
// 客户端来源
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetSource() string {
    return r.source
}
// Biz Setter
// 业务点
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

// Biz Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetBiz() string {
    return r.biz
}
// ExtraData Setter
// 额外信息
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetExtraData(extraData string) error {
    r.extraData = extraData
    r.Set("extra_data", extraData)
    return nil
}

// ExtraData Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetExtraData() string {
    return r.extraData
}
