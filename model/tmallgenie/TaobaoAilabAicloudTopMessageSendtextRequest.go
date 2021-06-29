package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
故事机发送文本留言 API请求
taobao.ailab.aicloud.top.message.sendtext

故事机文本留言
*/
type TaobaoAilabAicloudTopMessageSendtextRequest struct {
    model.Params
    // 账户体系隔离
    schema   string
    // 用户ID，此处传入第三方账户体系的用户 id
    userId   string
    // 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 扩展信息，用于存放APP类型等
    ext   string
    // 留言输入的文本
    text   string
}

// 初始化TaobaoAilabAicloudTopMessageSendtextRequest对象
func NewTaobaoAilabAicloudTopMessageSendtextRequest() *TaobaoAilabAicloudTopMessageSendtextRequest{
    return &TaobaoAilabAicloudTopMessageSendtextRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageSendtextRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.sendtext"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageSendtextRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageSendtextRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMessageSendtextRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户 id
func (r *TaobaoAilabAicloudTopMessageSendtextRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMessageSendtextRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageSendtextRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMessageSendtextRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageSendtextRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMessageSendtextRequest) GetExt() string {
    return r.ext
}
// Text Setter
// 留言输入的文本
func (r *TaobaoAilabAicloudTopMessageSendtextRequest) SetText(text string) error {
    r.text = text
    r.Set("text", text)
    return nil
}

// Text Getter
func (r TaobaoAilabAicloudTopMessageSendtextRequest) GetText() string {
    return r.text
}
