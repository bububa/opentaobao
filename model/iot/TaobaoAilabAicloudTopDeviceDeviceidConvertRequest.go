package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开放设备id转换内部设备id APIRequest
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

func NewTaobaoAilabAicloudTopDeviceDeviceidConvertRequest() *TaobaoAilabAicloudTopDeviceDeviceidConvertRequest{
    return &TaobaoAilabAicloudTopDeviceDeviceidConvertRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.deviceid.convert"
}

func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) SetDeviceOpenId(deviceOpenId string) error {
    r.deviceOpenId = deviceOpenId
    r.Set("device_open_id", deviceOpenId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetDeviceOpenId() string {
    return r.deviceOpenId
}

func (r *TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) SetSkillId(skillId string) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceDeviceidConvertRequest) GetSkillId() string {
    return r.skillId
}

