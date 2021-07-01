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
type AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest struct {
    model.Params
    // 用户OAuth授权后的token
    _token   string
    // 厂商在天猫精灵开放平台申请的技能id
    _skillId   string
    // 更新类型 1：设备更新 2：场景更新
    _type   string
}

// 初始化AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest对象
func NewAlibabaAilabsIotDeviceListUpdateNotifyRequest() *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest{
    return &AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetApiMethodName() string {
    return "alibaba.ailabs.iot.device.list.update.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 用户OAuth授权后的token
func (r *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetToken() string {
    return r._token
}
// SkillId Setter
// 厂商在天猫精灵开放平台申请的技能id
func (r *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) SetSkillId(_skillId string) error {
    r._skillId = _skillId
    r.Set("skill_id", _skillId)
    return nil
}

// SkillId Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetSkillId() string {
    return r._skillId
}
// Type Setter
// 更新类型 1：设备更新 2：场景更新
func (r *AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest) GetType() string {
    return r._type
}
