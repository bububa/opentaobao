package alink

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
消息提醒开关 APIRequest
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

func NewAlibabaAlinkMessageConfigSetRequest() *AlibabaAlinkMessageConfigSetRequest{
    return &AlibabaAlinkMessageConfigSetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlinkMessageConfigSetRequest) GetApiMethodName() string {
    return "alibaba.alink.message.config.set"
}

func (r AlibabaAlinkMessageConfigSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlinkMessageConfigSetRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaAlinkMessageConfigSetRequest) GetUuid() string {
    return r.uuid
}

func (r *AlibabaAlinkMessageConfigSetRequest) SetPushDisabled(pushDisabled string) error {
    r.pushDisabled = pushDisabled
    r.Set("push_disabled", pushDisabled)
    return nil
}

func (r AlibabaAlinkMessageConfigSetRequest) GetPushDisabled() string {
    return r.pushDisabled
}

