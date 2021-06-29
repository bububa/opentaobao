package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实人认证云服务提交接口 API请求
alibaba.security.jaq.rp.cloud.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpCloudSubmitRequest struct {
    model.Params
    // 认证token
    verifyToken   string
}

// 初始化AlibabaSecurityJaqRpCloudSubmitRequest对象
func NewAlibabaSecurityJaqRpCloudSubmitRequest() *AlibabaSecurityJaqRpCloudSubmitRequest{
    return &AlibabaSecurityJaqRpCloudSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudSubmitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpCloudSubmitRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudSubmitRequest) GetVerifyToken() string {
    return r.verifyToken
}
