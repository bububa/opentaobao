package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取openId设备授权码验证结果 API请求
taobao.ailab.aicloud.top.device.openid.authresult.get

获取openId设备授权码验证结果
*/
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest struct {
    model.Params
    // 淘宝openid
    _openId   string
    // 账户体系隔离
    _schema   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // authcode list
    _authCodes   []string
}

// 初始化TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceOpenidAuthresultGetRequest() *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest{
    return &TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.openid.authresult.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenId Setter
// 淘宝openid
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetOpenId() string {
    return r._openId
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetSchema() string {
    return r._schema
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetUtdId() string {
    return r._utdId
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetUserId() string {
    return r._userId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetExt() string {
    return r._ext
}
// AuthCodes Setter
// authcode list
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) SetAuthCodes(_authCodes []string) error {
    r._authCodes = _authCodes
    r.Set("auth_codes", _authCodes)
    return nil
}

// AuthCodes Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest) GetAuthCodes() []string {
    return r._authCodes
}
