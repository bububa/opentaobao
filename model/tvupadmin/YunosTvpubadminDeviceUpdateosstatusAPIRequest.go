package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceUpdateosstatusAPIRequest
更新系统版本审核状态 API请求
yunos.tvpubadmin.device.updateosstatus

更新系统版本审核状态 */
type YunosTvpubadminDeviceUpdateosstatusAPIRequest struct {
	model.Params
	// 升级ID
	_versionId int64
	// 牌照方
	_license int64
	// 审核状态
	_status string
	// 审核意见
	_auditComment string
}

// NewYunosTvpubadminDeviceUpdateosstatusRequest 初始化YunosTvpubadminDeviceUpdateosstatusAPIRequest对象
func NewYunosTvpubadminDeviceUpdateosstatusRequest() *YunosTvpubadminDeviceUpdateosstatusAPIRequest {
	return &YunosTvpubadminDeviceUpdateosstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.updateosstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is VersionId Setter
// 升级ID
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// Get VersionId Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// Set is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetLicense() int64 {
	return r._license
}

// Set is Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetStatus() string {
	return r._status
}

// Set is AuditComment Setter
// 审核意见
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// Get AuditComment Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetAuditComment() string {
	return r._auditComment
}
