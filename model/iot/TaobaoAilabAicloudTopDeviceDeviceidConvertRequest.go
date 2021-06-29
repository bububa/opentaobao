package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开放设备id转换内部设备id API请求
taobao.ailab.aicloud.top.device.deviceid.convert

将开放设备id转换为内部设备id
*/
type TaobaoAilabAicloudTopDeviceDeviceidConvertRequest struct {
    model.Params
    // 设备openId
    deviceOpenId   string
    // 技能id
    skillId   string
}

// 初始化TaobaoAilabAicloudTopDeviceDeviceidConvertRequest对象
func NewTaobaoAilabAicloudTopDeviceDeviceidConvertRequest() *TaobaoAilabAicloudTopDeviceDeviceidConvertRequest{
    return &TaobaoAilabAicloudTopDeviceDeviceidConvertRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.deviceid.convert"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceOpenId Setter
// 设备openId
func (r *TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) SetDeviceOpenId(deviceOpenId string) error {
    r.deviceOpenId = deviceOpenId
    r.Set("device_open_id", deviceOpenId)
    return nil
}

// DeviceOpenId Getter
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetDeviceOpenId() string {
    return r.deviceOpenId
}
// SkillId Setter
// 技能id
func (r *TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) SetSkillId(skillId string) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

// SkillId Getter
func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetSkillId() string {
    return r.skillId
}
