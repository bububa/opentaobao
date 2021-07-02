package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest 单视频审核提交审核结果 API请求
// yunos.tvpubadmin.content.single.video.submitauditresult
//
// 单视频审核提交审核结果
type YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest struct {
	model.Params
	// 审核人
	_auditor string
	// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
	_licenseState int64
	// 牌照方
	_license int64
	// 备注说明
	_auditComment string
	// 视频审核ID
	_videoAuditId int64
}

// NewYunosTvpubadminContentSingleVideoSubmitauditresultRequest 初始化YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest对象
func NewYunosTvpubadminContentSingleVideoSubmitauditresultRequest() *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest {
	return &YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.single.video.submitauditresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Auditor Setter
// 审核人
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetAuditor(_auditor string) error {
	r._auditor = _auditor
	r.Set("auditor", _auditor)
	return nil
}

// Get Auditor Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetAuditor() string {
	return r._auditor
}

// Set is LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetLicenseState(_licenseState int64) error {
	r._licenseState = _licenseState
	r.Set("license_state", _licenseState)
	return nil
}

// Get LicenseState Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetLicenseState() int64 {
	return r._licenseState
}

// Set is License Setter
// 牌照方
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetLicense() int64 {
	return r._license
}

// Set is AuditComment Setter
// 备注说明
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// Get AuditComment Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetAuditComment() string {
	return r._auditComment
}

// Set is VideoAuditId Setter
// 视频审核ID
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetVideoAuditId(_videoAuditId int64) error {
	r._videoAuditId = _videoAuditId
	r.Set("video_audit_id", _videoAuditId)
	return nil
}

// Get VideoAuditId Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetVideoAuditId() int64 {
	return r._videoAuditId
}
