package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceupdateosstatusAPIRequest 更新系统版本审核状态 API请求
// yunos.tvpubadmin.device.updateosstatus
//
// 更新系统版本审核状态
type YunostvpubadmindeviceupdateosstatusAPIRequest struct {
	model.Params
	// 审核状态
	_status string
	// 审核意见
	_auditComment string
	// 升级ID
	_versionId int64
	// 牌照方
	_license int64
}

// NewYunostvpubadmindeviceupdateosstatusRequest 初始化YunostvpubadmindeviceupdateosstatusAPIRequest对象
func NewYunostvpubadmindeviceupdateosstatusRequest() *YunostvpubadmindeviceupdateosstatusAPIRequest {
	return &YunostvpubadmindeviceupdateosstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceupdateosstatusAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.updateosstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceupdateosstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceupdateosstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 审核状态
func (r *YunostvpubadmindeviceupdateosstatusAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunostvpubadmindeviceupdateosstatusAPIRequest) GetStatus() string {
	return r._status
}

// SetAuditComment is AuditComment Setter
// 审核意见
func (r *YunostvpubadmindeviceupdateosstatusAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// GetAuditComment AuditComment Getter
func (r YunostvpubadmindeviceupdateosstatusAPIRequest) GetAuditComment() string {
	return r._auditComment
}

// SetVersionId is VersionId Setter
// 升级ID
func (r *YunostvpubadmindeviceupdateosstatusAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r YunostvpubadmindeviceupdateosstatusAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindeviceupdateosstatusAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindeviceupdateosstatusAPIRequest) GetLicense() int64 {
	return r._license
}
