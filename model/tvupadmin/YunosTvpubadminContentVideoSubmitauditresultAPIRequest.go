package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松提交视频审核结果 API请求
yunos.tvpubadmin.content.video.submitauditresult

迎客松提交视频审核结果
*/
type YunosTvpubadminContentVideoSubmitauditresultAPIRequest struct {
    model.Params
    // 审核结果，json格式
    _licenseAuditResult   string
}

// 初始化YunosTvpubadminContentVideoSubmitauditresultAPIRequest对象
func NewYunosTvpubadminContentVideoSubmitauditresultRequest() *YunosTvpubadminContentVideoSubmitauditresultAPIRequest{
    return &YunosTvpubadminContentVideoSubmitauditresultAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentVideoSubmitauditresultAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.video.submitauditresult"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentVideoSubmitauditresultAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LicenseAuditResult Setter
// 审核结果，json格式
func (r *YunosTvpubadminContentVideoSubmitauditresultAPIRequest) SetLicenseAuditResult(_licenseAuditResult string) error {
    r._licenseAuditResult = _licenseAuditResult
    r.Set("license_audit_result", _licenseAuditResult)
    return nil
}

// LicenseAuditResult Getter
func (r YunosTvpubadminContentVideoSubmitauditresultAPIRequest) GetLicenseAuditResult() string {
    return r._licenseAuditResult
}
