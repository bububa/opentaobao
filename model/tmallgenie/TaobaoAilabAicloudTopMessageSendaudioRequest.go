package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送语音留言 API请求
taobao.ailab.aicloud.top.message.sendaudio

将语音的二进制byte[]通过TOP接口发送保存
*/
type TaobaoAilabAicloudTopMessageSendaudioRequest struct {
    model.Params
    // 账户体系隔离
    schema   string
    // 用户ID，此处传入第三方账户体系的用户 id
    userId   string
    // 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 扩展信息，用于存放APP类型等
    ext   string
    // 语音的二进制
    message   []*model.File
}

// 初始化TaobaoAilabAicloudTopMessageSendaudioRequest对象
func NewTaobaoAilabAicloudTopMessageSendaudioRequest() *TaobaoAilabAicloudTopMessageSendaudioRequest{
    return &TaobaoAilabAicloudTopMessageSendaudioRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.sendaudio"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户 id
func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetExt() string {
    return r.ext
}
// Message Setter
// 语音的二进制
func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetMessage(message []*model.File) error {
    r.message = message
    r.Set("message", message)
    return nil
}

// Message Getter
func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetMessage() []*model.File {
    return r.message
}
