package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmindeviceappupgradedetailAPIRequest 获取应用升级详情 API请求
// yunos.tvpubadmin.device.appupgradedetail
//
// 获取应用升级详情
type YunostvpubadmindeviceappupgradedetailAPIRequest struct {
	model.Params
	// 应用升级的ID
	_versionId int64
	// 牌照方
	_license int64
}

// NewYunostvpubadmindeviceappupgradedetailRequest 初始化YunostvpubadmindeviceappupgradedetailAPIRequest对象
func NewYunostvpubadmindeviceappupgradedetailRequest() *YunostvpubadmindeviceappupgradedetailAPIRequest {
	return &YunostvpubadmindeviceappupgradedetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmindeviceappupgradedetailAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.device.appupgradedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmindeviceappupgradedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmindeviceappupgradedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVersionId is VersionId Setter
// 应用升级的ID
func (r *YunostvpubadmindeviceappupgradedetailAPIRequest) SetVersionId(_versionId int64) error {
	r._versionId = _versionId
	r.Set("version_id", _versionId)
	return nil
}

// GetVersionId VersionId Getter
func (r YunostvpubadmindeviceappupgradedetailAPIRequest) GetVersionId() int64 {
	return r._versionId
}

// SetLicense is License Setter
// 牌照方
func (r *YunostvpubadmindeviceappupgradedetailAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmindeviceappupgradedetailAPIRequest) GetLicense() int64 {
	return r._license
}
