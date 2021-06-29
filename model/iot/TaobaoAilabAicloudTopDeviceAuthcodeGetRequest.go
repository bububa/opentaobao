package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码 API请求
taobao.ailab.aicloud.top.device.authcode.get

获取设备授权码
*/
type TaobaoAilabAicloudTopDeviceAuthcodeGetRequest struct {
    model.Params
    // 账户体系隔离，即硬件接入平台中取得的schema key。
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id，由开发者或厂商自行定义，每一schema key下唯一即可
    _userId   string
    // (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
}

// 初始化TaobaoAilabAicloudTopDeviceAuthcodeGetRequest对象
func NewTaobaoAilabAicloudTopDeviceAuthcodeGetRequest() *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest{
    return &TaobaoAilabAicloudTopDeviceAuthcodeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.authcode.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离，即硬件接入平台中取得的schema key。
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id，由开发者或厂商自行定义，每一schema key下唯一即可
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetRequest) GetExt() string {
    return r._ext
}
