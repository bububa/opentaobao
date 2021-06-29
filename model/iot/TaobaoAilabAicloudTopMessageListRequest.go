package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取留言列表 API请求
taobao.ailab.aicloud.top.message.list

根据指定参数获取留言列表
*/
type TaobaoAilabAicloudTopMessageListRequest struct {
    model.Params
    // 账户体系隔离
    schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 扩展信息，用于存放APP类型等
    ext   string
    // 截止时间，yyyy-MM-dd HH:mm:ss.SSS
    beforeTime   string
    // 一次性展示多少条信息
    limit   int64
}

// 初始化TaobaoAilabAicloudTopMessageListRequest对象
func NewTaobaoAilabAicloudTopMessageListRequest() *TaobaoAilabAicloudTopMessageListRequest{
    return &TaobaoAilabAicloudTopMessageListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageListRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMessageListRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopMessageListRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMessageListRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageListRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMessageListRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageListRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMessageListRequest) GetExt() string {
    return r.ext
}
// BeforeTime Setter
// 截止时间，yyyy-MM-dd HH:mm:ss.SSS
func (r *TaobaoAilabAicloudTopMessageListRequest) SetBeforeTime(beforeTime string) error {
    r.beforeTime = beforeTime
    r.Set("before_time", beforeTime)
    return nil
}

// BeforeTime Getter
func (r TaobaoAilabAicloudTopMessageListRequest) GetBeforeTime() string {
    return r.beforeTime
}
// Limit Setter
// 一次性展示多少条信息
func (r *TaobaoAilabAicloudTopMessageListRequest) SetLimit(limit int64) error {
    r.limit = limit
    r.Set("limit", limit)
    return nil
}

// Limit Getter
func (r TaobaoAilabAicloudTopMessageListRequest) GetLimit() int64 {
    return r.limit
}
