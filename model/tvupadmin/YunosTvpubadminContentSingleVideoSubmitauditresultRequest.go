package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单视频审核提交审核结果 API请求
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

// 初始化YunosTvpubadminContentSingleVideoSubmitauditresultRequest对象
func NewYunosTvpubadminContentSingleVideoSubmitauditresultRequest() *YunosTvpubadminContentSingleVideoSubmitauditresultRequest{
    return &YunosTvpubadminContentSingleVideoSubmitauditresultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.content.single.video.submitauditresult"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Auditor Setter
// 审核人
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetAuditor(auditor string) error {
    r.auditor = auditor
    r.Set("auditor", auditor)
    return nil
}

// Auditor Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetAuditor() string {
    return r.auditor
}
// LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetLicenseState(licenseState int64) error {
    r.licenseState = licenseState
    r.Set("license_state", licenseState)
    return nil
}

// LicenseState Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetLicenseState() int64 {
    return r.licenseState
}
// License Setter
// 牌照方
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetLicense(license int64) error {
    r.license = license
    r.Set("license", license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetLicense() int64 {
    return r.license
}
// AuditComment Setter
// 备注说明
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetAuditComment(auditComment string) error {
    r.auditComment = auditComment
    r.Set("audit_comment", auditComment)
    return nil
}

// AuditComment Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetAuditComment() string {
    return r.auditComment
}
// VideoAuditId Setter
// 视频审核ID
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetVideoAuditId(videoAuditId int64) error {
    r.videoAuditId = videoAuditId
    r.Set("video_audit_id", videoAuditId)
    return nil
}

// VideoAuditId Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetVideoAuditId() int64 {
    return r.videoAuditId
}
