package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.osupgradedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is VersionId Setter
// 版本ID
func (r *YunosTvpubadminDeviceOsupgradedetailAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// Get VersionId Getter
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// Set is License Setter
// 牌照方
func (r *YunosTvpubadminDeviceOsupgradedetailAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// Get License Getter
func (r YunosTvpubadminDeviceOsupgradedetailAPIRequest) GetLicense() int64 {
	return r._license
}
