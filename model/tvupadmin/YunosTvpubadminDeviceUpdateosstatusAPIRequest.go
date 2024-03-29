package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceUpdateosstatusAPIRequest 更新系统版本审核状态 API请求
// yunos.tvpubadmin.device.updateosstatus
//
// 更新系统版本审核状态
type YunosTvpubadminDeviceUpdateosstatusAPIRequest struct {
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

// NewYunosTvpubadminDeviceUpdateosstatusRequest 初始化YunosTvpubadminDeviceUpdateosstatusAPIRequest对象
func NewYunosTvpubadminDeviceUpdateosstatusRequest() *YunosTvpubadminDeviceUpdateosstatusAPIRequest {
	return &YunosTvpubadminDeviceUpdateosstatusAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) Reset() {
	r._status = ""
	r._auditComment = ""
	r._versionId = 0
	r._license = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.updateosstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 审核状态
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetStatus() string {
	return r._status
}

// SetAuditComment is AuditComment Setter
// 审核意见
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetAuditComment(_auditComment string) error {
	r._auditComment = _auditComment
	r.Set("audit_comment", _auditComment)
	return nil
}

// GetAuditComment AuditComment Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetAuditComment() string {
	return r._auditComment
}

// SetVersionId is VersionId Setter
// 升级ID
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceUpdateosstatusAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceUpdateosstatusAPIRequest) GetLicense() int64 {
	return r._license
}

var poolYunosTvpubadminDeviceUpdateosstatusAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceUpdateosstatusRequest()
	},
}

// GetYunosTvpubadminDeviceUpdateosstatusRequest 从 sync.Pool 获取 YunosTvpubadminDeviceUpdateosstatusAPIRequest
func GetYunosTvpubadminDeviceUpdateosstatusAPIRequest() *YunosTvpubadminDeviceUpdateosstatusAPIRequest {
	return poolYunosTvpubadminDeviceUpdateosstatusAPIRequest.Get().(*YunosTvpubadminDeviceUpdateosstatusAPIRequest)
}

// ReleaseYunosTvpubadminDeviceUpdateosstatusAPIRequest 将 YunosTvpubadminDeviceUpdateosstatusAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceUpdateosstatusAPIRequest(v *YunosTvpubadminDeviceUpdateosstatusAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceUpdateosstatusAPIRequest.Put(v)
}
