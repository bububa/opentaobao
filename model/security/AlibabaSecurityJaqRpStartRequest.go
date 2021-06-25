package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证开始 APIRequest
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

func NewAlibabaSecurityJaqRpStartRequest() *AlibabaSecurityJaqRpStartRequest{
    return &AlibabaSecurityJaqRpStartRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpStartRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.start"
}

func (r AlibabaSecurityJaqRpStartRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpStartRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpStartRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpStartRequest) SetClientInfo(clientInfo *RpClientInfo) error {
    r.clientInfo = clientInfo
    r.Set("client_info", clientInfo)
    return nil
}

func (r AlibabaSecurityJaqRpStartRequest) GetClientInfo() *RpClientInfo {
    return r.clientInfo
}

func (r *AlibabaSecurityJaqRpStartRequest) SetExtraData(extraData string) error {
    r.extraData = extraData
    r.Set("extra_data", extraData)
    return nil
}

func (r AlibabaSecurityJaqRpStartRequest) GetExtraData() string {
    return r.extraData
}

