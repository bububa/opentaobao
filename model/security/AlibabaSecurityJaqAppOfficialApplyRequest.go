package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全官方应用申请 APIRequest
alibaba.security.jaq.app.official.apply

官方应用申请接口
*/
type AlibabaSecurityJaqAppOfficialApplyRequest struct {
    model.Params

    // 官方应用申请入参
    officialAppApplyRequest   *OfficialAppApplyRequest 

}

func NewAlibabaSecurityJaqAppOfficialApplyRequest() *AlibabaSecurityJaqAppOfficialApplyRequest{
    return &AlibabaSecurityJaqAppOfficialApplyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppOfficialApplyRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.official.apply"
}

func (r AlibabaSecurityJaqAppOfficialApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppOfficialApplyRequest) SetOfficialAppApplyRequest(officialAppApplyRequest *OfficialAppApplyRequest) error {
    r.officialAppApplyRequest = officialAppApplyRequest
    r.Set("official_app_apply_request", officialAppApplyRequest)
    return nil
}

func (r AlibabaSecurityJaqAppOfficialApplyRequest) GetOfficialAppApplyRequest() *OfficialAppApplyRequest {
    return r.officialAppApplyRequest
}

