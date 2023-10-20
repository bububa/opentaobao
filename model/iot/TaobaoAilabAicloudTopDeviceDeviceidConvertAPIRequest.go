package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest 开放设备id转换内部设备id API请求
// taobao.ailab.aicloud.top.device.deviceid.convert
//
// 将开放设备id转换为内部设备id
type TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest struct {
	model.Params
	// 设备openId
	_deviceOpenId string
	// 技能id
	_skillId string
}

// NewTaobaoAilabAicloudTopDeviceDeviceidConvertRequest 初始化TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceDeviceidConvertRequest() *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest {
	return &TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) Reset() {
	r._deviceOpenId = ""
	r._skillId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.deviceid.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceOpenId is DeviceOpenId Setter
// 设备openId
func (r *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) SetDeviceOpenId(_deviceOpenId string) error {
	r._deviceOpenId = _deviceOpenId
	r.Set("device_open_id", _deviceOpenId)
	return nil
}

// GetDeviceOpenId DeviceOpenId Getter
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) GetDeviceOpenId() string {
	return r._deviceOpenId
}

// SetSkillId is SkillId Setter
// 技能id
func (r *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) SetSkillId(_skillId string) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) GetSkillId() string {
	return r._skillId
}

var poolTaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceDeviceidConvertRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceDeviceidConvertRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest
func GetTaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest() *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest 将 TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest(v *TaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceDeviceidConvertAPIRequest.Put(v)
}
