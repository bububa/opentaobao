package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openid设备通用授权码 API请求
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

// 初始化TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest对象
func NewTaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest() *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest{
    return &TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.openid.authcode.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenId Setter
// 淘宝openid
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

// OpenId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetOpenId() string {
    return r.openId
}
// Schema Setter
// 账户体系隔离，即硬件接入平台中取得的schema key。
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetSchema() string {
    return r.schema
}
// UtdId Setter
// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest) GetExt() string {
    return r.ext
}
