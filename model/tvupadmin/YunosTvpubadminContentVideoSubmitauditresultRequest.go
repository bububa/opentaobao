package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
迎客松提交视频审核结果 APIRequest
yunos.tvpubadmin.content.video.submitauditresult

迎客松提交视频审核结果
*/
type YunosTvpubadminContentVideoSubmitauditresultRequest struct {
    model.Params

    // 审核结果，json格式
    licenseAuditResult   string 

}

func NewYunosTvpubadminContentVideoSubmitauditresultRequest() *YunosTvpubadminContentVideoSubmitauditresultRequest{
    return &YunosTvpubadminContentVideoSubmitauditresultRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentVideoSubmitauditresultRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.video.submitauditresult"
}

func (r YunosTvpubadminContentVideoSubmitauditresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentVideoSubmitauditresultRequest) SetLicenseAuditResult(licenseAuditResult string) error {
    r.licenseAuditResult = licenseAuditResult
    r.Set("license_audit_result", licenseAuditResult)
    return nil
}

func (r YunosTvpubadminContentVideoSubmitauditresultRequest) GetLicenseAuditResult() string {
    return r.licenseAuditResult
}

