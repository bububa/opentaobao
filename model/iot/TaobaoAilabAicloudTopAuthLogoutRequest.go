package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
登出 API请求
taobao.ailab.aicloud.top.auth.logout

登出
*/
type TaobaoAilabAicloudTopAuthLogoutAPIRequest struct {
    model.Params
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 账户体系隔离，建议传入设备uuid
    _schema   string
}

// 初始化TaobaoAilabAicloudTopAuthLogoutAPIRequest对象
func NewTaobaoAilabAicloudTopAuthLogoutRequest() *TaobaoAilabAicloudTopAuthLogoutAPIRequest{
    return &TaobaoAilabAicloudTopAuthLogoutAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopAuthLogoutAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.auth.logout"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopAuthLogoutAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopAuthLogoutAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopAuthLogoutAPIRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopAuthLogoutAPIRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopAuthLogoutAPIRequest) GetUtdId() string {
    return r._utdId
}
// Schema Setter
// 账户体系隔离，建议传入设备uuid
func (r *TaobaoAilabAicloudTopAuthLogoutAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopAuthLogoutAPIRequest) GetSchema() string {
    return r._schema
}
