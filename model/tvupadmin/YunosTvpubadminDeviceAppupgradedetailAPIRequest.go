package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminDeviceAppupgradedetailAPIRequest 获取应用升级详情 API请求
// yunos.tvpubadmin.device.appupgradedetail
//
// 获取应用升级详情
type YunosTvpubadminDeviceAppupgradedetailAPIRequest struct {
	model.Params
	// 应用升级的ID
	_versionId int64
	// 牌照方
	_license int64
}

// NewYunosTvpubadminDeviceAppupgradedetailRequest 初始化YunosTvpubadminDeviceAppupgradedetailAPIRequest对象
func NewYunosTvpubadminDeviceAppupgradedetailRequest() *YunosTvpubadminDeviceAppupgradedetailAPIRequest {
	return &YunosTvpubadminDeviceAppupgradedetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminDeviceAppupgradedetailAPIRequest) Reset() {
	r._versionId = 0
	r._license = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.appupgradedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVersionId is VersionId Setter
// 应用升级的ID
func (r *YunosTvpubadminDeviceAppupgradedetailAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetLicense is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceAppupgradedetailAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminDeviceAppupgradedetailAPIRequest) GetLicense() int64 {
	return r._license
}

var poolYunosTvpubadminDeviceAppupgradedetailAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminDeviceAppupgradedetailRequest()
	},
}

// GetYunosTvpubadminDeviceAppupgradedetailRequest 从 sync.Pool 获取 YunosTvpubadminDeviceAppupgradedetailAPIRequest
func GetYunosTvpubadminDeviceAppupgradedetailAPIRequest() *YunosTvpubadminDeviceAppupgradedetailAPIRequest {
	return poolYunosTvpubadminDeviceAppupgradedetailAPIRequest.Get().(*YunosTvpubadminDeviceAppupgradedetailAPIRequest)
}

// ReleaseYunosTvpubadminDeviceAppupgradedetailAPIRequest 将 YunosTvpubadminDeviceAppupgradedetailAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminDeviceAppupgradedetailAPIRequest(v *YunosTvpubadminDeviceAppupgradedetailAPIRequest) {
	v.Reset()
	poolYunosTvpubadminDeviceAppupgradedetailAPIRequest.Put(v)
}
