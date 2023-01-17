package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkRexoutResourceListCheckAPIRequest ReX应用中心资源更新检测-外部 API请求
// wdk.rexout.resource.list.check
//
// ReX应用中心资源更新检测-外部
type WdkRexoutResourceListCheckAPIRequest struct {
	model.Params
	// 主包名
	_packageName string
	// 设备ID
	_deviceId string
	// 设备类型
	_deviceType string
	// 旧版本资源信息
	_oldVersions string
	// 所属商家结构
	_orgInfo string
	// 自定义参数
	_attributes string
	// 主包版本
	_versionCode int64
	// 租户ID
	_tenantId int64
}

// NewWdkRexoutResourceListCheckRequest 初始化WdkRexoutResourceListCheckAPIRequest对象
func NewWdkRexoutResourceListCheckRequest() *WdkRexoutResourceListCheckAPIRequest {
	return &WdkRexoutResourceListCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkRexoutResourceListCheckAPIRequest) GetApiMethodName() string {
	return "wdk.rexout.resource.list.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkRexoutResourceListCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkRexoutResourceListCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPackageName is PackageName Setter
// 主包名
func (r *WdkRexoutResourceListCheckAPIRequest) SetPackageName(_packageName string) error {
	r._packageName = _packageName
	r.Set("package_name", _packageName)
	return nil
}

// GetPackageName PackageName Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetPackageName() string {
	return r._packageName
}

// SetDeviceId is DeviceId Setter
// 设备ID
func (r *WdkRexoutResourceListCheckAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetDeviceType is DeviceType Setter
// 设备类型
func (r *WdkRexoutResourceListCheckAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// GetDeviceType DeviceType Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// SetOldVersions is OldVersions Setter
// 旧版本资源信息
func (r *WdkRexoutResourceListCheckAPIRequest) SetOldVersions(_oldVersions string) error {
	r._oldVersions = _oldVersions
	r.Set("old_versions", _oldVersions)
	return nil
}

// GetOldVersions OldVersions Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetOldVersions() string {
	return r._oldVersions
}

// SetOrgInfo is OrgInfo Setter
// 所属商家结构
func (r *WdkRexoutResourceListCheckAPIRequest) SetOrgInfo(_orgInfo string) error {
	r._orgInfo = _orgInfo
	r.Set("org_info", _orgInfo)
	return nil
}

// GetOrgInfo OrgInfo Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetOrgInfo() string {
	return r._orgInfo
}

// SetAttributes is Attributes Setter
// 自定义参数
func (r *WdkRexoutResourceListCheckAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetVersionCode is VersionCode Setter
// 主包版本
func (r *WdkRexoutResourceListCheckAPIRequest) SetVersionCode(_versionCode int64) error {
	r._versionCode = _versionCode
	r.Set("version_code", _versionCode)
	return nil
}

// GetVersionCode VersionCode Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetVersionCode() int64 {
	return r._versionCode
}

// SetTenantId is TenantId Setter
// 租户ID
func (r *WdkRexoutResourceListCheckAPIRequest) SetTenantId(_tenantId int64) error {
	r._tenantId = _tenantId
	r.Set("tenant_id", _tenantId)
	return nil
}

// GetTenantId TenantId Getter
func (r WdkRexoutResourceListCheckAPIRequest) GetTenantId() int64 {
	return r._tenantId
}
