package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取认证会话token APIRequest
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

func NewAlibabaSecurityJaqRpGetverifytokenRequest() *AlibabaSecurityJaqRpGetverifytokenRequest{
    return &AlibabaSecurityJaqRpGetverifytokenRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.getverifytoken"
}

func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetAccountId(accountId string) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetAccountId() string {
    return r.accountId
}

func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetTicketId(ticketId string) error {
    r.ticketId = ticketId
    r.Set("ticket_id", ticketId)
    return nil
}

func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetTicketId() string {
    return r.ticketId
}

func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetSource() string {
    return r.source
}

func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetBiz(biz string) error {
    r.biz = biz
    r.Set("biz", biz)
    return nil
}

func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetBiz() string {
    return r.biz
}

func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetExtraData(extraData string) error {
    r.extraData = extraData
    r.Set("extra_data", extraData)
    return nil
}

func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetExtraData() string {
    return r.extraData
}

