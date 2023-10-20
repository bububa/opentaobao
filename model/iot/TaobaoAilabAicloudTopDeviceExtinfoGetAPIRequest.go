package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest 获取设备扩展信息 API请求
// taobao.ailab.aicloud.top.device.extinfo.get
//
// 获取设备扩展信息
type TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest struct {
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

// NewTaobaoAilabAicloudTopDeviceExtinfoGetRequest 初始化TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceExtinfoGetRequest() *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest {
	return &TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) Reset() {
	r._originUserId = ""
	r._schemaKey = ""
	r._userType = ""
	r._deviceId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.extinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginUserId is OriginUserId Setter
// 三方id、淘宝openId
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetOriginUserId(_originUserId string) error {
	r._originUserId = _originUserId
	r.Set("origin_user_id", _originUserId)
	return nil
}

// GetOriginUserId OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetOriginUserId() string {
	return r._originUserId
}

// SetSchemaKey is SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetUserType is UserType Setter
// 类型：openTaoBao, extUser
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetUserType() string {
	return r._userType
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

var poolTaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceExtinfoGetRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceExtinfoGetRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest
func GetTaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest() *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest 将 TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest(v *TaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceExtinfoGetAPIRequest.Put(v)
}
