package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云开始认证 APIRequest
alibaba.security.jaq.rp.cloud.start

聚安全实人认证开始
*/
type AlibabaSecurityJaqRpCloudStartRequest struct {
    model.Params

    // token
    verifyToken   string 

    // 客户端信息，如果是服务端接入，里面的参数可为空
    clientInfo   *RpClientInfo 

    // 扩展信息
    extraData   string 

}

func NewAlibabaSecurityJaqRpCloudStartRequest() *AlibabaSecurityJaqRpCloudStartRequest{
    return &AlibabaSecurityJaqRpCloudStartRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpCloudStartRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.start"
}

func (r AlibabaSecurityJaqRpCloudStartRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpCloudStartRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpCloudStartRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpCloudStartRequest) SetClientInfo(clientInfo *RpClientInfo) error {
    r.clientInfo = clientInfo
    r.Set("client_info", clientInfo)
    return nil
}

func (r AlibabaSecurityJaqRpCloudStartRequest) GetClientInfo() *RpClientInfo {
    return r.clientInfo
}

func (r *AlibabaSecurityJaqRpCloudStartRequest) SetExtraData(extraData string) error {
    r.extraData = extraData
    r.Set("extra_data", extraData)
    return nil
}

func (r AlibabaSecurityJaqRpCloudStartRequest) GetExtraData() string {
    return r.extraData
}

