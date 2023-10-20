package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceupdateappstatusAPIRequest 更新应用版本审核状态 API请求
// yunos.tvpubadmin.device.updateappstatus
//
// 更新应用版本审核状态
type YunostvpubadmindeviceupdateappstatusAPIRequest struct {
	model.Params
	// 审核状态
	_status string
	// 审核意见
	_auditComment string
	// 应用ID
	_versionId int64
	// 牌照方
	_license int64
}

// NewYunostvpubadmindeviceupdateappstatusRequest 初始化YunostvpubadmindeviceupdateappstatusAPIRequest对象
func NewYunostvpubadmindeviceupdateappstatusRequest() *YunostvpubadmindeviceupdateappstatusAPIRequest {
	return &YunostvpubadmindeviceupdateappstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceupdateappstatusAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.updateappstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceupdateappstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceupdateappstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 审核状态
func (r *YunostvpubadmindeviceupdateappstatusAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunostvpubadmindeviceupdateappstatusAPIRequest) GetStatus() string {
	return r._status
}

// SetAuditComment is AuditComment Setter
// 审核意见
func (r *YunostvpubadmindeviceupdateappstatusAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// GetAuditComment AuditComment Getter
func (r YunostvpubadmindeviceupdateappstatusAPIRequest) GetAuditComment() string {
	return r._auditComment
}

// SetVersionId is VersionId Setter
// 应用ID
func (r *YunostvpubadmindeviceupdateappstatusAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r YunostvpubadmindeviceupdateappstatusAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindeviceupdateappstatusAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindeviceupdateappstatusAPIRequest) GetLicense() int64 {
	return r._license
}
