package alilabs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备列表更新通知 API请求
alibaba.ailabs.iot.device.list.update.notify

用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
*/
type AlibabaAilabsIotDeviceListUpdateNotifyRequest struct {
    model.Params
    // 用户OAuth授权后的token
    token   string
    // 厂商在天猫精灵开放平台申请的技能id
    skillId   string
    // 更新类型 1：设备更新 2：场景更新
    type   string
}

// 初始化AlibabaAilabsIotDeviceListUpdateNotifyRequest对象
func NewAlibabaAilabsIotDeviceListUpdateNotifyRequest() *AlibabaAilabsIotDeviceListUpdateNotifyRequest{
    return &AlibabaAilabsIotDeviceListUpdateNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceListUpdateNotifyRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.device.list.update.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceListUpdateNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 用户OAuth授权后的token
func (r *AlibabaAilabsIotDeviceListUpdateNotifyRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyRequest) GetToken() string {
    return r.token
}
// SkillId Setter
// 厂商在天猫精灵开放平台申请的技能id
func (r *AlibabaAilabsIotDeviceListUpdateNotifyRequest) SetSkillId(skillId string) error {
    r.skillId = skillId
    r.Set("skill_id", skillId)
    return nil
}

// SkillId Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyRequest) GetSkillId() string {
    return r.skillId
}
// Type Setter
// 更新类型 1：设备更新 2：场景更新
func (r *AlibabaAilabsIotDeviceListUpdateNotifyRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyRequest) GetType() string {
    return r.type
}
