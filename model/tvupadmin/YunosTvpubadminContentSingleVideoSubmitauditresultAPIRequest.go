package tvupadmin

import (
	"net/url"
	"sync"

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
	// 备注说明
	_auditComment string
	// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
	_licenseState int64
	// 牌照方
	_license int64
	// 视频审核ID
	_videoAuditId int64
}

// NewYunosTvpubadminContentSingleVideoSubmitauditresultRequest 初始化YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest对象
func NewYunosTvpubadminContentSingleVideoSubmitauditresultRequest() *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest {
	return &YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) Reset() {
	r._auditor = ""
	r._auditComment = ""
	r._licenseState = 0
	r._license = 0
	r._videoAuditId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.single.video.submitauditresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuditor is Auditor Setter
// 审核人
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetAuditor(_auditor string) error {
	r._auditor = _auditor
	r.Set("auditor", _auditor)
	return nil
}

// GetAuditor Auditor Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetAuditor() string {
	return r._auditor
}

// SetAuditComment is AuditComment Setter
// 备注说明
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// GetAuditComment AuditComment Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetAuditComment() string {
	return r._auditComment
}

// SetLicenseState is LicenseState Setter
// 审核状态：1未提审，2审核中，3通过，4不通过，5已下线
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetLicenseState(_licenseState int64) error {
	r._licenseState = _licenseState
	r.Set("license_state", _licenseState)
	return nil
}

// GetLicenseState LicenseState Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetLicenseState() int64 {
	return r._licenseState
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetLicense() int64 {
	return r._license
}

// SetVideoAuditId is VideoAuditId Setter
// 视频审核ID
func (r *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) SetVideoAuditId(_videoAuditId int64) error {
	r._videoAuditId = _videoAuditId
	r.Set("video_audit_id", _videoAuditId)
	return nil
}

// GetVideoAuditId VideoAuditId Getter
func (r YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) GetVideoAuditId() int64 {
	return r._videoAuditId
}

var poolYunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentSingleVideoSubmitauditresultRequest()
	},
}

// GetYunosTvpubadminContentSingleVideoSubmitauditresultRequest 从 sync.Pool 获取 YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest
func GetYunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest() *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest {
	return poolYunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest.Get().(*YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest)
}

// ReleaseYunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest 将 YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest(v *YunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentSingleVideoSubmitauditresultAPIRequest.Put(v)
}
