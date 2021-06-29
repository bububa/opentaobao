package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全实人认证获取结果接口 API请求
alibaba.security.jaq.rp.fetchmaterial

聚安全实人认证获取结果接口
*/
type AlibabaSecurityJaqRpFetchmaterialRequest struct {
    model.Params
    // 消息服务推送的key
    _securityKey   string
}

// 初始化AlibabaSecurityJaqRpFetchmaterialRequest对象
func NewAlibabaSecurityJaqRpFetchmaterialRequest() *AlibabaSecurityJaqRpFetchmaterialRequest{
    return &AlibabaSecurityJaqRpFetchmaterialRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpFetchmaterialRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.fetchmaterial"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpFetchmaterialRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SecurityKey Setter
// 消息服务推送的key
func (r *AlibabaSecurityJaqRpFetchmaterialRequest) SetSecurityKey(_securityKey string) error {
    r._securityKey = _securityKey
    r.Set("security_key", _securityKey)
    return nil
}

// SecurityKey Getter
func (r AlibabaSecurityJaqRpFetchmaterialRequest) GetSecurityKey() string {
    return r._securityKey
}
