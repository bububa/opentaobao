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
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // authCodes信息
    _authCodes   []string
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
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetExt() string {
    return r._ext
}
// AuthCodes Setter
// authCodes信息
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetRequest) SetAuthCodes(_authCodes []string) error {
    r._authCodes = _authCodes
    r.Set("auth_codes", _authCodes)
    return nil
}

// AuthCodes Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetRequest) GetAuthCodes() []string {
    return r._authCodes
}
