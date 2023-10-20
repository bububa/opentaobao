package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest 开放设备id转换内部设备id API请求
// taobao.ailab.aicloud.top.device.deviceid.convert
//
// 将开放设备id转换为内部设备id
type TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest struct {
	model.Params
	// 设备openId
	_deviceOpenId string
	// 技能id
	_skillId string
}

// NewTaobaoailabaicloudtopdevicedeviceidconvertRequest 初始化TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest对象
func NewTaobaoailabaicloudtopdevicedeviceidconvertRequest() *TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest {
	return &TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.deviceid.convert"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceOpenId is DeviceOpenId Setter
// 设备openId
func (r *TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest) SetDeviceOpenId(_deviceOpenId string) error {
	r._deviceOpenId = _deviceOpenId
	r.Set("device_open_id", _deviceOpenId)
	return nil
}

// GetDeviceOpenId DeviceOpenId Getter
func (r TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest) GetDeviceOpenId() string {
	return r._deviceOpenId
}

// SetSkillId is SkillId Setter
// 技能id
func (r *TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest) SetSkillId(_skillId string) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest) GetSkillId() string {
	return r._skillId
}
