package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openId设备授权码验证结果 APIRequest
taobao.ailab.aicloud.top.device.openid.authresult.get

获取openId设备授权码验证结果
*/
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest struct {
    model.Params

    // 淘宝openid
    openId   string 

    // 账户体系隔离
    schema   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

    // authcode list
    authCodes   []string 

}

func NewTaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest() *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest{
    return &TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.openid.authresult.get"
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetOpenId() string {
    return r.openId
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) SetAuthCodes(authCodes []string) error {
    r.authCodes = authCodes
    r.Set("auth_codes", authCodes)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest) GetAuthCodes() []string {
    return r.authCodes
}

