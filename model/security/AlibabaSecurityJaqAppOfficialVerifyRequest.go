package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全验证官方应用接口 APIRequest
alibaba.security.jaq.app.official.verify

接入用户来查询应用是否为官方应用
*/
type AlibabaSecurityJaqAppOfficialVerifyRequest struct {
    model.Params

    // 验证参数
    officialAppVerifyRequest   *OfficialAppVerifyRequest 

}

func NewAlibabaSecurityJaqAppOfficialVerifyRequest() *AlibabaSecurityJaqAppOfficialVerifyRequest{
    return &AlibabaSecurityJaqAppOfficialVerifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppOfficialVerifyRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.official.verify"
}

func (r AlibabaSecurityJaqAppOfficialVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppOfficialVerifyRequest) SetOfficialAppVerifyRequest(officialAppVerifyRequest *OfficialAppVerifyRequest) error {
    r.officialAppVerifyRequest = officialAppVerifyRequest
    r.Set("official_app_verify_request", officialAppVerifyRequest)
    return nil
}

func (r AlibabaSecurityJaqAppOfficialVerifyRequest) GetOfficialAppVerifyRequest() *OfficialAppVerifyRequest {
    return r.officialAppVerifyRequest
}

