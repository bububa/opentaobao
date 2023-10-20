package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceosupgradedetailAPIRequest 获取系统升级详情 API请求
// yunos.tvpubadmin.device.osupgradedetail
//
// 获取系统升级详情
type YunostvpubadmindeviceosupgradedetailAPIRequest struct {
	model.Params
	// 版本ID
	_versionId int64
	// 牌照方
	_license int64
}

// NewYunostvpubadmindeviceosupgradedetailRequest 初始化YunostvpubadmindeviceosupgradedetailAPIRequest对象
func NewYunostvpubadmindeviceosupgradedetailRequest() *YunostvpubadmindeviceosupgradedetailAPIRequest {
	return &YunostvpubadmindeviceosupgradedetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceosupgradedetailAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.osupgradedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceosupgradedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceosupgradedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVersionId is VersionId Setter
// 版本ID
func (r *YunostvpubadmindeviceosupgradedetailAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r YunostvpubadmindeviceosupgradedetailAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindeviceosupgradedetailAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindeviceosupgradedetailAPIRequest) GetLicense() int64 {
	return r._license
}
