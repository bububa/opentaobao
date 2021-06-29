package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送留言 API请求
taobao.ailab.aicloud.top.message.send

供准入的外部用户实现发送留言功能，APP端发送，设备端读取
*/
type TaobaoAilabAicloudTopMessageSendRequest struct {
    model.Params
    // 扩展信息，用于存放APP类型等
    ext   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 账户体系隔离
    schema   string
    // 用户上传到OSS上的文件路径，不包含域名
    audioPath   string
}

// 初始化TaobaoAilabAicloudTopMessageSendRequest对象
func NewTaobaoAilabAicloudTopMessageSendRequest() *TaobaoAilabAicloudTopMessageSendRequest{
    return &TaobaoAilabAicloudTopMessageSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageSendRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageSendRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMessageSendRequest) GetExt() string {
    return r.ext
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageSendRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMessageSendRequest) GetUtdId() string {
    return r.utdId
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopMessageSendRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMessageSendRequest) GetUserId() string {
    return r.userId
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageSendRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMessageSendRequest) GetSchema() string {
    return r.schema
}
// AudioPath Setter
// 用户上传到OSS上的文件路径，不包含域名
func (r *TaobaoAilabAicloudTopMessageSendRequest) SetAudioPath(audioPath string) error {
    r.audioPath = audioPath
    r.Set("audio_path", audioPath)
    return nil
}

// AudioPath Getter
func (r TaobaoAilabAicloudTopMessageSendRequest) GetAudioPath() string {
    return r.audioPath
}
