package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openid设备通用授权码 APIRequest
taobao.ailab.aicloud.top.device.openid.authcode.get

获取openid设备通用授权码
*/
type TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest struct {
    model.Params

    // 淘宝openid
    openId   string 

    // 账户体系隔离，即硬件接入平台中取得的schema key。
    schema   string 

    // (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

}

func NewTaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest() *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest{
    return &TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.openid.authcode.get"
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetOpenId() string {
    return r.openId
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetExt() string {
    return r.ext
}

