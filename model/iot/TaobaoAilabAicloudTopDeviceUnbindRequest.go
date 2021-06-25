package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 APIRequest
taobao.ailab.aicloud.top.device.unbind

解绑设备
*/
type TaobaoAilabAicloudTopDeviceUnbindRequest struct {
    model.Params

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

}

func NewTaobaoAilabAicloudTopDeviceUnbindRequest() *TaobaoAilabAicloudTopDeviceUnbindRequest{
    return &TaobaoAilabAicloudTopDeviceUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceUnbindRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.unbind"
}

func (r TaobaoAilabAicloudTopDeviceUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceUnbindRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceUnbindRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopDeviceUnbindRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceUnbindRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopDeviceUnbindRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceUnbindRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopDeviceUnbindRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceUnbindRequest) GetExt() string {
    return r.ext
}

