package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicestatusinfogetAPIRequest 获取设备状态信息 API请求
// taobao.ailab.aicloud.top.device.statusinfo.get
//
// 获取设备状态信息
type TaobaoailabaicloudtopdevicestatusinfogetAPIRequest struct {
	model.Params
	// 三方用户id或淘宝openId
	_originUserId string
	// 账号秘钥
	_schemaKey string
	// 三方传extUser，淘宝传openTaoBaoUser
	_userType string
	// 设备id
	_deviceId string
}

// NewTaobaoailabaicloudtopdevicestatusinfogetRequest 初始化TaobaoailabaicloudtopdevicestatusinfogetAPIRequest对象
func NewTaobaoailabaicloudtopdevicestatusinfogetRequest() *TaobaoailabaicloudtopdevicestatusinfogetAPIRequest {
	return &TaobaoailabaicloudtopdevicestatusinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.statusinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginUserId is OriginUserId Setter
// 三方用户id或淘宝openId
func (r *TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) SetOriginUserId(_originUserId string) error {
	r._originUserId = _originUserId
	r.Set("origin_user_id", _originUserId)
	return nil
}

// GetOriginUserId OriginUserId Getter
func (r TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) GetOriginUserId() string {
	return r._originUserId
}

// SetSchemaKey is SchemaKey Setter
// 账号秘钥
func (r *TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetUserType is UserType Setter
// 三方传extUser，淘宝传openTaoBaoUser
func (r *TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) GetUserType() string {
	return r._userType
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoailabaicloudtopdevicestatusinfogetAPIRequest) GetDeviceId() string {
	return r._deviceId
}
