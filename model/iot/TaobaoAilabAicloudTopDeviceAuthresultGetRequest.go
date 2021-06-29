package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码验证结果 API请求
taobao.ailab.aicloud.top.device.authresult.get

获取设备授权码验证结果
*/
type TaobaoAilabAicloudTopDeviceAuthresultGetRequest struct {
    model.Params
    // 账户体系隔离
    schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 扩展信息，用于存放APP类型等
    ext   string
    // authCodes信息
    authCodes   []string
}

// 初始化TaobaoAilabAicloudTopDeviceAuthresultGetRequest对象
func NewTaobaoAilabAicloudTopDeviceAuthresultGetRequest() *TaobaoAilabAicloudTopDeviceAuthresultGetRequest{
    return &TaobaoAilabAicloudTopDeviceAuthresultGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.authresult.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetExt() string {
    return r.ext
}
// AuthCodes Setter
// authCodes信息
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetAuthCodes(authCodes []string) error {
    r.authCodes = authCodes
    r.Set("auth_codes", authCodes)
    return nil
}

// AuthCodes Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetAuthCodes() []string {
    return r.authCodes
}
