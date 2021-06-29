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
    verifyToken   string
    // 客户端信息，如果是服务端接入，里面的参数可为空
    clientInfo   *RpClientInfo
    // 扩展信息
    extraData   string
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
func (r *AlibabaSecurityJaqRpStartRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpStartRequest) GetVerifyToken() string {
    return r.verifyToken
}
// ClientInfo Setter
// 客户端信息，如果是服务端接入，里面的参数可为空
func (r *AlibabaSecurityJaqRpStartRequest) SetClientInfo(clientInfo *RpClientInfo) error {
    r.clientInfo = clientInfo
    r.Set("client_info", clientInfo)
    return nil
}

// ClientInfo Getter
func (r AlibabaSecurityJaqRpStartRequest) GetClientInfo() *RpClientInfo {
    return r.clientInfo
}
// ExtraData Setter
// 扩展信息
func (r *AlibabaSecurityJaqRpStartRequest) SetExtraData(extraData string) error {
    r.extraData = extraData
    r.Set("extra_data", extraData)
    return nil
}

// ExtraData Getter
func (r AlibabaSecurityJaqRpStartRequest) GetExtraData() string {
    return r.extraData
}
