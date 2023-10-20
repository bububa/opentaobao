package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceextinfogetAPIRequest 获取设备扩展信息 API请求
// taobao.ailab.aicloud.top.device.extinfo.get
//
// 获取设备扩展信息
type TaobaoailabaicloudtopdeviceextinfogetAPIRequest struct {
	model.Params
	// 三方id、淘宝openId
	_originUserId string
	// 账号秘钥
	_schemaKey string
	// 类型：openTaoBao, extUser
	_userType string
	// 设备id
	_deviceId string
}

// NewTaobaoailabaicloudtopdeviceextinfogetRequest 初始化TaobaoailabaicloudtopdeviceextinfogetAPIRequest对象
func NewTaobaoailabaicloudtopdeviceextinfogetRequest() *TaobaoailabaicloudtopdeviceextinfogetAPIRequest {
	return &TaobaoailabaicloudtopdeviceextinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdeviceextinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.extinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdeviceextinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdeviceextinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginUserId is OriginUserId Setter
// 三方id、淘宝openId
func (r *TaobaoailabaicloudtopdeviceextinfogetAPIRequest) SetOriginUserId(_originUserId string) error {
	r._originUserId = _originUserId
	r.Set("origin_user_id", _originUserId)
	return nil
}

// GetOriginUserId OriginUserId Getter
func (r TaobaoailabaicloudtopdeviceextinfogetAPIRequest) GetOriginUserId() string {
	return r._originUserId
}

// SetSchemaKey is SchemaKey Setter
// 账号秘钥
func (r *TaobaoailabaicloudtopdeviceextinfogetAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r TaobaoailabaicloudtopdeviceextinfogetAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetUserType is UserType Setter
// 类型：openTaoBao, extUser
func (r *TaobaoailabaicloudtopdeviceextinfogetAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r TaobaoailabaicloudtopdeviceextinfogetAPIRequest) GetUserType() string {
	return r._userType
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoailabaicloudtopdeviceextinfogetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoailabaicloudtopdeviceextinfogetAPIRequest) GetDeviceId() string {
	return r._deviceId
}
