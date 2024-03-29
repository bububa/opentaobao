package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest 获取设备详细信息 API请求
// taobao.ailab.aicloud.top.device.detailinfo.get
//
// 获取设备详细信息
type TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest struct {
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

// NewTaobaoAilabAicloudTopDeviceDetailinfoGetRequest 初始化TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceDetailinfoGetRequest() *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest {
	return &TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) Reset() {
	r._originUserId = ""
	r._schemaKey = ""
	r._userType = ""
	r._deviceId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.detailinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginUserId is OriginUserId Setter
// 三方用户id或淘宝openId
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetOriginUserId(_originUserId string) error {
	r._originUserId = _originUserId
	r.Set("origin_user_id", _originUserId)
	return nil
}

// GetOriginUserId OriginUserId Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetOriginUserId() string {
	return r._originUserId
}

// SetSchemaKey is SchemaKey Setter
// 账号秘钥
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetSchemaKey(_schemaKey string) error {
	r._schemaKey = _schemaKey
	r.Set("schema_key", _schemaKey)
	return nil
}

// GetSchemaKey SchemaKey Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetSchemaKey() string {
	return r._schemaKey
}

// SetUserType is UserType Setter
// 三方传extUser，淘宝传openTaoBaoUser
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetUserType() string {
	return r._userType
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

var poolTaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceDetailinfoGetRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceDetailinfoGetRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest
func GetTaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest() *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest 将 TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest(v *TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest.Put(v)
}
