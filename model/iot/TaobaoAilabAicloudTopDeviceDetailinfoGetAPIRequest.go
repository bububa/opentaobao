package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicedetailinfogetAPIRequest 获取设备详细信息 API请求
// taobao.ailab.aicloud.top.device.detailinfo.get
//
// 获取设备详细信息
type TaobaoailabaicloudtopdevicedetailinfogetAPIRequest struct {
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

// NewTaobaoailabaicloudtopdevicedetailinfogetRequest 初始化TaobaoailabaicloudtopdevicedetailinfogetAPIRequest对象
func NewTaobaoailabaicloudtopdevicedetailinfogetRequest() *TaobaoailabaicloudtopdevicedetailinfogetAPIRequest {
	return &TaobaoailabaicloudtopdevicedetailinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.detailinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginUserId is OriginUserId Setter
// 三方用户id或淘宝openId
func (r *TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) SetOriginUserId(_originUserId string) error {
	r._originUserId = _originUserId
	r.Set("origin_user_id", _originUserId)
	return nil
}

// GetOriginUserId OriginUserId Getter
func (r TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) GetOriginUserId() string {
	return r._originUserId
}

// SetSchemaKey is SchemaKey Setter
// 账号秘钥
func (r *TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetUserType is UserType Setter
// 三方传extUser，淘宝传openTaoBaoUser
func (r *TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) GetUserType() string {
	return r._userType
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoailabaicloudtopdevicedetailinfogetAPIRequest) GetDeviceId() string {
	return r._deviceId
}
