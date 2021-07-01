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
type AlibabaSecurityJaqRpSubmitAPIRequest struct {
    model.Params
    // 认证token
    _verifyToken   string
}

// 初始化AlibabaSecurityJaqRpSubmitAPIRequest对象
func NewAlibabaSecurityJaqRpSubmitRequest() *AlibabaSecurityJaqRpSubmitAPIRequest{
    return &AlibabaSecurityJaqRpSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpSubmitAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpSubmitAPIRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpSubmitAPIRequest) GetVerifyToken() string {
    return r._verifyToken
}
