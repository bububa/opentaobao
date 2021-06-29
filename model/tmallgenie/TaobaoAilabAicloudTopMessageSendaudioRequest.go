package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送语音留言 APIRequest
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

func NewTaobaoAilabAicloudTopMessageSendaudioRequest() *TaobaoAilabAicloudTopMessageSendaudioRequest{
    return &TaobaoAilabAicloudTopMessageSendaudioRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.sendaudio"
}

func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopMessageSendaudioRequest) SetMessage(message []*model.File) error {
    r.message = message
    r.Set("message", message)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendaudioRequest) GetMessage() []*model.File {
    return r.message
}

