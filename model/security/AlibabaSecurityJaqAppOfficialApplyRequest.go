package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚安全官方应用申请 API请求
alibaba.security.jaq.app.official.apply

官方应用申请接口
*/
type AlibabaSecurityJaqAppOfficialApplyAPIRequest struct {
    model.Params
    // 官方应用申请入参
    _officialAppApplyRequest   *OfficialAppApplyRequest
}

// 初始化AlibabaSecurityJaqAppOfficialApplyAPIRequest对象
func NewAlibabaSecurityJaqAppOfficialApplyRequest() *AlibabaSecurityJaqAppOfficialApplyAPIRequest{
    return &AlibabaSecurityJaqAppOfficialApplyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqAppOfficialApplyAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.official.apply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqAppOfficialApplyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfficialAppApplyRequest Setter
// 官方应用申请入参
func (r *AlibabaSecurityJaqAppOfficialApplyAPIRequest) SetOfficialAppApplyRequest(_officialAppApplyRequest *OfficialAppApplyRequest) error {
    r._officialAppApplyRequest = _officialAppApplyRequest
    r.Set("official_app_apply_request", _officialAppApplyRequest)
    return nil
}

// OfficialAppApplyRequest Getter
func (r AlibabaSecurityJaqAppOfficialApplyAPIRequest) GetOfficialAppApplyRequest() *OfficialAppApplyRequest {
    return r._officialAppApplyRequest
}
