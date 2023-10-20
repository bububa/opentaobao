package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowdkiotdeviceadminmqtttokengetAPIRequest 获取MQTT访问令牌 API请求
// taobao.wdk.iot.deviceadmin.mqtt.token.get
//
// 智能硬件设备动态注册和获取mqtt设备信息
type TaobaowdkiotdeviceadminmqtttokengetAPIRequest struct {
	model.Params
	// accessKey
	_accessKey string
	// 申请令牌的客户端时间戳
	_applyTimestamp int64
}

// NewTaobaowdkiotdeviceadminmqtttokengetRequest 初始化TaobaowdkiotdeviceadminmqtttokengetAPIRequest对象
func NewTaobaowdkiotdeviceadminmqtttokengetRequest() *TaobaowdkiotdeviceadminmqtttokengetAPIRequest {
	return &TaobaowdkiotdeviceadminmqtttokengetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowdkiotdeviceadminmqtttokengetAPIRequest) GetApiMethodName() string {
	return "taobao.wdk.iot.deviceadmin.mqtt.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowdkiotdeviceadminmqtttokengetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowdkiotdeviceadminmqtttokengetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// accessKey
func (r *TaobaowdkiotdeviceadminmqtttokengetAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r TaobaowdkiotdeviceadminmqtttokengetAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetApplyTimestamp is ApplyTimestamp Setter
// 申请令牌的客户端时间戳
func (r *TaobaowdkiotdeviceadminmqtttokengetAPIRequest) SetApplyTimestamp(_applyTimestamp int64) error {
	r._applyTimestamp = _applyTimestamp
	r.Set("apply_timestamp", _applyTimestamp)
	return nil
}

// GetApplyTimestamp ApplyTimestamp Getter
func (r TaobaowdkiotdeviceadminmqtttokengetAPIRequest) GetApplyTimestamp() int64 {
	return r._applyTimestamp
}
