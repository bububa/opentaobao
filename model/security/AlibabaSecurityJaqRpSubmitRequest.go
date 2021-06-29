package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证提交认证接口 API请求
alibaba.security.jaq.rp.submit

聚安全实人认证提交认证接口
*/
type AlibabaSecurityJaqRpSubmitRequest struct {
    model.Params
    // 认证token
    _verifyToken   string
}

// 初始化AlibabaSecurityJaqRpSubmitRequest对象
func NewAlibabaSecurityJaqRpSubmitRequest() *AlibabaSecurityJaqRpSubmitRequest{
    return &AlibabaSecurityJaqRpSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpSubmitRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpSubmitRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpSubmitRequest) GetVerifyToken() string {
    return r._verifyToken
}
