package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceOsupgradedetailAPIRequest 获取系统升级详情 API请求
// yunos.tvpubadmin.device.osupgradedetail
//
// 获取系统升级详情
type YunosTvpubadminDeviceOsupgradedetailAPIRequest struct {
	model.Params
	// 版本ID
	_versionId int64
	// 牌照方
	_license int64
}

// NewYunosTvpubadminDeviceOsupgradedetailRequest 初始化YunosTvpubadminDeviceOsupgradedetailAPIRequest对象
func NewYunosTvpubadminDeviceOsupgradedetailRequest() *YunosTvpubadminDeviceOsupgradedetailAPIRequest {
	return &YunosTvpubadminDeviceOsupgradedetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceOsupgradedetailAPIRequest) Reset() {
	r._versionId = 0
	r._license = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.osupgradedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVersionId is VersionId Setter
// 版本ID
func (r *YunosTvpubadminDeviceOsupgradedetailAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceOsupgradedetailAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetLicense() int64 {
	return r._license
}

var poolYunosTvpubadminDeviceOsupgradedetailAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceOsupgradedetailRequest()
	},
}

// GetYunosTvpubadminDeviceOsupgradedetailRequest 从 sync.Pool 获取 YunosTvpubadminDeviceOsupgradedetailAPIRequest
func GetYunosTvpubadminDeviceOsupgradedetailAPIRequest() *YunosTvpubadminDeviceOsupgradedetailAPIRequest {
	return poolYunosTvpubadminDeviceOsupgradedetailAPIRequest.Get().(*YunosTvpubadminDeviceOsupgradedetailAPIRequest)
}

// ReleaseYunosTvpubadminDeviceOsupgradedetailAPIRequest 将 YunosTvpubadminDeviceOsupgradedetailAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceOsupgradedetailAPIRequest(v *YunosTvpubadminDeviceOsupgradedetailAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceOsupgradedetailAPIRequest.Put(v)
}
