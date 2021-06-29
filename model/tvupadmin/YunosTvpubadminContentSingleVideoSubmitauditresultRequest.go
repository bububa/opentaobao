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
    _auditor   string
    // 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
    _licenseState   int64
    // 牌照方
    _license   int64
    // 备注说明
    _auditComment   string
    // 视频审核ID
    _videoAuditId   int64
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
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetAuditor(_auditor string) error {
    r._auditor = _auditor
    r.Set("auditor", _auditor)
    return nil
}

// Auditor Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetAuditor() string {
    return r._auditor
}
// LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetLicenseState(_licenseState int64) error {
    r._licenseState = _licenseState
    r.Set("license_state", _licenseState)
    return nil
}

// LicenseState Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetLicenseState() int64 {
    return r._licenseState
}
// License Setter
// 牌照方
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetLicense(_license int64) error {
    r._license = _license
    r.Set("license", _license)
    return nil
}

// License Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetLicense() int64 {
    return r._license
}
// AuditComment Setter
// 备注说明
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetAuditComment(_auditComment string) error {
    r._auditComment = _auditComment
    r.Set("audit_comment", _auditComment)
    return nil
}

// AuditComment Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetAuditComment() string {
    return r._auditComment
}
// VideoAuditId Setter
// 视频审核ID
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultRequest) SetVideoAuditId(_videoAuditId int64) error {
    r._videoAuditId = _videoAuditId
    r.Set("video_audit_id", _videoAuditId)
    return nil
}

// VideoAuditId Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultRequest) GetVideoAuditId() int64 {
    return r._videoAuditId
}
