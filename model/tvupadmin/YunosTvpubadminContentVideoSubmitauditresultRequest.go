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
type YunosTvpubadminContentVideoSubmitauditresultRequest struct {
    model.Params
    // 审核结果，json格式
    licenseAuditResult   string
}

// 初始化YunosTvpubadminContentVideoSubmitauditresultRequest对象
func NewYunosTvpubadminContentVideoSubmitauditresultRequest() *YunosTvpubadminContentVideoSubmitauditresultRequest{
    return &YunosTvpubadminContentVideoSubmitauditresultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentVideoSubmitauditresultRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.video.submitauditresult"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentVideoSubmitauditresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LicenseAuditResult Setter
// 审核结果，json格式
func (r *YunosTvpubadminContentVideoSubmitauditresultRequest) SetLicenseAuditResult(licenseAuditResult string) error {
    r.licenseAuditResult = licenseAuditResult
    r.Set("license_audit_result", licenseAuditResult)
    return nil
}

// LicenseAuditResult Getter
func (r YunosTvpubadminContentVideoSubmitauditresultRequest) GetLicenseAuditResult() string {
    return r.licenseAuditResult
}
