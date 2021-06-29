package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证查询认证结果 API请求
alibaba.security.jaq.rp.query

聚安全实人认证查询认证结果
*/
type AlibabaSecurityJaqRpQueryRequest struct {
    model.Params
    // token
    _verifyToken   string
}

// 初始化AlibabaSecurityJaqRpQueryRequest对象
func NewAlibabaSecurityJaqRpQueryRequest() *AlibabaSecurityJaqRpQueryRequest{
    return &AlibabaSecurityJaqRpQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpQueryRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// token
func (r *AlibabaSecurityJaqRpQueryRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpQueryRequest) GetVerifyToken() string {
    return r._verifyToken
}
