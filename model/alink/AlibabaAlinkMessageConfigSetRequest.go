package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息提醒开关 API请求
alibaba.alink.message.config.set

阿里智能消息开关
*/
type AlibabaAlinkMessageConfigSetRequest struct {
    model.Params
    // 设备id
    uuid   string
    // 0：开启，1：关闭
    pushDisabled   string
}

// 初始化AlibabaAlinkMessageConfigSetRequest对象
func NewAlibabaAlinkMessageConfigSetRequest() *AlibabaAlinkMessageConfigSetRequest{
    return &AlibabaAlinkMessageConfigSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageConfigSetRequest) GetApiMethodName() string {
    return "alibaba.alink.message.config.set"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageConfigSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkMessageConfigSetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkMessageConfigSetRequest) GetUuid() string {
    return r.uuid
}
// PushDisabled Setter
// 0：开启，1：关闭
func (r *AlibabaAlinkMessageConfigSetRequest) SetPushDisabled(pushDisabled string) error {
    r.pushDisabled = pushDisabled
    r.Set("push_disabled", pushDisabled)
    return nil
}

// PushDisabled Getter
func (r AlibabaAlinkMessageConfigSetRequest) GetPushDisabled() string {
    return r.pushDisabled
}
