package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发送留言 APIRequest
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

func NewTaobaoAilabAicloudTopMessageSendRequest() *TaobaoAilabAicloudTopMessageSendRequest{
    return &TaobaoAilabAicloudTopMessageSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMessageSendRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.send"
}

func (r TaobaoAilabAicloudTopMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMessageSendRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopMessageSendRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopMessageSendRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopMessageSendRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopMessageSendRequest) SetAudioPath(audioPath string) error {
    r.audioPath = audioPath
    r.Set("audio_path", audioPath)
    return nil
}

func (r TaobaoAilabAicloudTopMessageSendRequest) GetAudioPath() string {
    return r.audioPath
}

