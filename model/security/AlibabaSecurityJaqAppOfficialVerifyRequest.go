package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全验证官方应用接口 API请求
alibaba.security.jaq.app.official.verify

接入用户来查询应用是否为官方应用
*/
type AlibabaSecurityJaqAppOfficialVerifyRequest struct {
    model.Params
    // 验证参数
    _officialAppVerifyRequest   *OfficialAppVerifyRequest
}

// 初始化AlibabaSecurityJaqAppOfficialVerifyRequest对象
func NewAlibabaSecurityJaqAppOfficialVerifyRequest() *AlibabaSecurityJaqAppOfficialVerifyRequest{
    return &AlibabaSecurityJaqAppOfficialVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppOfficialVerifyRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.official.verify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppOfficialVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAppVerifyRequest Setter
// 验证参数
func (r *AlibabaSecurityJaqAppOfficialVerifyRequest) SetOfficialAppVerifyRequest(_officialAppVerifyRequest *OfficialAppVerifyRequest) error {
    r._officialAppVerifyRequest = _officialAppVerifyRequest
    r.Set("official_app_verify_request", _officialAppVerifyRequest)
    return nil
}

// OfficialAppVerifyRequest Getter
func (r AlibabaSecurityJaqAppOfficialVerifyRequest) GetOfficialAppVerifyRequest() *OfficialAppVerifyRequest {
    return r._officialAppVerifyRequest
}
