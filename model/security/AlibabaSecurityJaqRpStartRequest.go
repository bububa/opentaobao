package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证开始 API请求
alibaba.security.jaq.rp.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpStartRequest struct {
    model.Params
    // token
    _verifyToken   string
    // 客户端信息，如果是服务端接入，里面的参数可为空
    _clientInfo   *RpClientInfo
    // 扩展信息
    _extraData   string
}

// 初始化AlibabaSecurityJaqRpStartRequest对象
func NewAlibabaSecurityJaqRpStartRequest() *AlibabaSecurityJaqRpStartRequest{
    return &AlibabaSecurityJaqRpStartRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpStartRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.start"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpStartRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpStartRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpStartRequest) GetVerifyToken() string {
    return r._verifyToken
}
// ClientInfo Setter
// 客户端信息，如果是服务端接入，里面的参数可为空
func (r *AlibabaSecurityJaqRpStartRequest) SetClientInfo(_clientInfo *RpClientInfo) error {
    r._clientInfo = _clientInfo
    r.Set("client_info", _clientInfo)
    return nil
}

// ClientInfo Getter
func (r AlibabaSecurityJaqRpStartRequest) GetClientInfo() *RpClientInfo {
    return r._clientInfo
}
// ExtraData Setter
// 扩展信息
func (r *AlibabaSecurityJaqRpStartRequest) SetExtraData(_extraData string) error {
    r._extraData = _extraData
    r.Set("extra_data", _extraData)
    return nil
}

// ExtraData Getter
func (r AlibabaSecurityJaqRpStartRequest) GetExtraData() string {
    return r._extraData
}
