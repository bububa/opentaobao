package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单视频审核提交审核结果 APIRequest
yunos.tvpubadmin.content.single.video.submitauditresult

单视频审核提交审核结果
*/
type YunosTvpubadminContentSingleVideoSubmitauditresultRequest struct {
    model.Params

    // 审核人
    auditor   string 

    // 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
    licenseState   int64 

    // 牌照方
    license   int64 

    // 备注说明
    auditComment   string 

    // 视频审核ID
    videoAuditId   int64 

}

func NewYunosTvpubadminContentSingleVideoSubmitauditresultRequest() *YunosTvpubadminContentSingleVideoSubmitauditresultRequest{
    return &YunosTvpubadminContentSingleVideoSubmitauditresultRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.single.video.submitauditresult"
}

func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetAuditor(auditor string) error {
    r.auditor = auditor
    r.Set("auditor", auditor)
    return nil
}

func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetAuditor() string {
    return r.auditor
}

func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetLicenseState(licenseState int64) error {
    r.licenseState = licenseState
    r.Set("license_state", licenseState)
    return nil
}

func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetLicenseState() int64 {
    return r.licenseState
}

func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetLicense() int64 {
    return r.license
}

func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetAuditComment(auditComment string) error {
    r.auditComment = auditComment
    r.Set("audit_comment", auditComment)
    return nil
}

func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetAuditComment() string {
    return r.auditComment
}

func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetVideoAuditId(videoAuditId int64) error {
    r.videoAuditId = videoAuditId
    r.Set("video_audit_id", videoAuditId)
    return nil
}

func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetVideoAuditId() int64 {
    return r.videoAuditId
}

