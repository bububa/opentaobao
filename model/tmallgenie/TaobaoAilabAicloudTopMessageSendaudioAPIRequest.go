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
type TaobaoAilabAicloudTopMessageSendaudioAPIRequest struct {
    model.Params
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户 id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 语音的二进制
    _message   *model.File
}

// 初始化TaobaoAilabAicloudTopMessageSendaudioAPIRequest对象
func NewTaobaoAilabAicloudTopMessageSendaudioRequest() *TaobaoAilabAicloudTopMessageSendaudioAPIRequest{
    return &TaobaoAilabAicloudTopMessageSendaudioAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageSendaudioAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.sendaudio"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageSendaudioAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageSendaudioAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMessageSendaudioAPIRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户 id
func (r *TaobaoAilabAicloudTopMessageSendaudioAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMessageSendaudioAPIRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageSendaudioAPIRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMessageSendaudioAPIRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageSendaudioAPIRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMessageSendaudioAPIRequest) GetExt() string {
    return r._ext
}
// Message Setter
// 语音的二进制
func (r *TaobaoAilabAicloudTopMessageSendaudioAPIRequest) SetMessage(_message *model.File) error {
    r._message = _message
    r.Set("message", _message)
    return nil
}

// Message Getter
func (r TaobaoAilabAicloudTopMessageSendaudioAPIRequest) GetMessage() *model.File {
    return r._message
}
