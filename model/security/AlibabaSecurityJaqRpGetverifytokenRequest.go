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
    _accountId   string
    // 选填，作为一次验证的唯一标识，每次验证需更换。如果不是验证类型可不填
    _ticketId   string
    // 客户端来源
    _source   string
    // 业务点
    _biz   string
    // 额外信息
    _extraData   string
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
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetAccountId(_accountId string) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetAccountId() string {
    return r._accountId
}
// TicketId Setter
// 选填，作为一次验证的唯一标识，每次验证需更换。如果不是验证类型可不填
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetTicketId(_ticketId string) error {
    r._ticketId = _ticketId
    r.Set("ticket_id", _ticketId)
    return nil
}

// TicketId Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetTicketId() string {
    return r._ticketId
}
// Source Setter
// 客户端来源
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetSource() string {
    return r._source
}
// Biz Setter
// 业务点
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetBiz(_biz string) error {
    r._biz = _biz
    r.Set("biz", _biz)
    return nil
}

// Biz Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetBiz() string {
    return r._biz
}
// ExtraData Setter
// 额外信息
func (r *AlibabaSecurityJaqRpGetverifytokenRequest) SetExtraData(_extraData string) error {
    r._extraData = _extraData
    r.Set("extra_data", _extraData)
    return nil
}

// ExtraData Getter
func (r AlibabaSecurityJaqRpGetverifytokenRequest) GetExtraData() string {
    return r._extraData
}
