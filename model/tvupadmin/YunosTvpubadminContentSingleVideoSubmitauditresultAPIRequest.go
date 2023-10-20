package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest 单视频审核提交审核结果 API请求
// yunos.tvpubadmin.content.single.video.submitauditresult
//
// 单视频审核提交审核结果
type YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest struct {
	model.Params
	// 审核人
	_auditor string
	// 备注说明
	_auditComment string
	// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
	_licenseState int64
	// 牌照方
	_license int64
	// 视频审核ID
	_videoAuditId int64
}

// NewYunostvpubadmincontentsinglevideosubmitauditresultRequest 初始化YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest对象
func NewYunostvpubadmincontentsinglevideosubmitauditresultRequest() *YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest {
	return &YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.single.video.submitauditresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuditor is Auditor Setter
// 审核人
func (r *YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) SetAuditor(_auditor string) error {
	r._auditor = _auditor
	r.Set("auditor", _auditor)
	return nil
}

// GetAuditor Auditor Getter
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetAuditor() string {
	return r._auditor
}

// SetAuditComment is AuditComment Setter
// 备注说明
func (r *YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// GetAuditComment AuditComment Getter
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetAuditComment() string {
	return r._auditComment
}

// SetLicenseState is LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) SetLicenseState(_licenseState int64) error {
	r._licenseState = _licenseState
	r.Set("license_state", _licenseState)
	return nil
}

// GetLicenseState LicenseState Getter
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetLicenseState() int64 {
	return r._licenseState
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetLicense() int64 {
	return r._license
}

// SetVideoAuditId is VideoAuditId Setter
// 视频审核ID
func (r *YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) SetVideoAuditId(_videoAuditId int64) error {
	r._videoAuditId = _videoAuditId
	r.Set("video_audit_id", _videoAuditId)
	return nil
}

// GetVideoAuditId VideoAuditId Getter
func (r YunostvpubadmincontentsinglevideosubmitauditresultAPIRequest) GetVideoAuditId() int64 {
	return r._videoAuditId
}
