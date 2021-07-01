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
type AlibabaAlinkMessageConfigSetAPIRequest struct {
    model.Params
    // 设备id
    _uuid   string
    // 0：开启，1：关闭
    _pushDisabled   string
}

// 初始化AlibabaAlinkMessageConfigSetAPIRequest对象
func NewAlibabaAlinkMessageConfigSetRequest() *AlibabaAlinkMessageConfigSetAPIRequest{
    return &AlibabaAlinkMessageConfigSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetApiMethodName() string {
    return "alibaba.alink.message.config.set"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *AlibabaAlinkMessageConfigSetAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetUuid() string {
    return r._uuid
}
// PushDisabled Setter
// 0：开启，1：关闭
func (r *AlibabaAlinkMessageConfigSetAPIRequest) SetPushDisabled(_pushDisabled string) error {
    r._pushDisabled = _pushDisabled
    r.Set("push_disabled", _pushDisabled)
    return nil
}

// PushDisabled Getter
func (r AlibabaAlinkMessageConfigSetAPIRequest) GetPushDisabled() string {
    return r._pushDisabled
}
