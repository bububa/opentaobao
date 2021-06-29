package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
登陆 API请求
taobao.ailab.aicloud.top.auth.get

登陆
*/
type TaobaoAilabAicloudTopAuthGetRequest struct {
    model.Params
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 账户体系隔离
    _schema   string
    // app类型
    _appType   string
}

// 初始化TaobaoAilabAicloudTopAuthGetRequest对象
func NewTaobaoAilabAicloudTopAuthGetRequest() *TaobaoAilabAicloudTopAuthGetRequest{
    return &TaobaoAilabAicloudTopAuthGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopAuthGetRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.auth.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopAuthGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopAuthGetRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopAuthGetRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopAuthGetRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopAuthGetRequest) GetUtdId() string {
    return r._utdId
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopAuthGetRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopAuthGetRequest) GetSchema() string {
    return r._schema
}
// AppType Setter
// app类型
func (r *TaobaoAilabAicloudTopAuthGetRequest) SetAppType(_appType string) error {
    r._appType = _appType
    r.Set("app_type", _appType)
    return nil
}

// AppType Getter
func (r TaobaoAilabAicloudTopAuthGetRequest) GetAppType() string {
    return r._appType
}
