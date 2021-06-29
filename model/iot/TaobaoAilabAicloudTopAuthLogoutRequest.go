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
type TaobaoAilabAicloudTopAuthLogoutRequest struct {
    model.Params
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 账户体系隔离，建议传入设备uuid
    schema   string
}

// 初始化TaobaoAilabAicloudTopAuthLogoutRequest对象
func NewTaobaoAilabAicloudTopAuthLogoutRequest() *TaobaoAilabAicloudTopAuthLogoutRequest{
    return &TaobaoAilabAicloudTopAuthLogoutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopAuthLogoutRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.auth.logout"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopAuthLogoutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopAuthLogoutRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopAuthLogoutRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopAuthLogoutRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopAuthLogoutRequest) GetUtdId() string {
    return r.utdId
}
// Schema Setter
// 账户体系隔离，建议传入设备uuid
func (r *TaobaoAilabAicloudTopAuthLogoutRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopAuthLogoutRequest) GetSchema() string {
    return r.schema
}
