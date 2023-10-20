package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest 获取MQTT访问令牌 API请求
// taobao.wdk.iot.deviceadmin.mqtt.token.get
//
// 智能硬件设备动态注册和获取mqtt设备信息
type TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest struct {
	model.Params
	// accessKey
	_accessKey string
	// 申请令牌的客户端时间戳
	_applyTimestamp int64
}

// NewTaobaoWdkIotDeviceadminMqttTokenGetRequest 初始化TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest对象
func NewTaobaoWdkIotDeviceadminMqttTokenGetRequest() *TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest {
	return &TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) Reset() {
	r._accessKey = ""
	r._applyTimestamp = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.iot.deviceadmin.mqtt.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// accessKey
func (r *TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetApplyTimestamp is ApplyTimestamp Setter
// 申请令牌的客户端时间戳
func (r *TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) SetApplyTimestamp(_applyTimestamp int64) error {
	r._applyTimestamp = _applyTimestamp
	r.Set("apply_timestamp", _applyTimestamp)
	return nil
}

// GetApplyTimestamp ApplyTimestamp Getter
func (r TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) GetApplyTimestamp() int64 {
	return r._applyTimestamp
}

var poolTaobaoWdkIotDeviceadminMqttTokenGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWdkIotDeviceadminMqttTokenGetRequest()
	},
}

// GetTaobaoWdkIotDeviceadminMqttTokenGetRequest 从 sync.Pool 获取 TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest
func GetTaobaoWdkIotDeviceadminMqttTokenGetAPIRequest() *TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest {
	return poolTaobaoWdkIotDeviceadminMqttTokenGetAPIRequest.Get().(*TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest)
}

// ReleaseTaobaoWdkIotDeviceadminMqttTokenGetAPIRequest 将 TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWdkIotDeviceadminMqttTokenGetAPIRequest(v *TaobaoWdkIotDeviceadminMqttTokenGetAPIRequest) {
	v.Reset()
	poolTaobaoWdkIotDeviceadminMqttTokenGetAPIRequest.Put(v)
}
