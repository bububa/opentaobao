package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单条对话流信息 APIRequest
taobao.ailab.aicloud.top.feedlist.delete

删除指定的某一条对话流信息
*/
type TaobaoAilabAicloudTopFeedlistDeleteRequest struct {
    model.Params

    // 扩展信息，用于存放APP类型等
    ext   string 

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 消息的唯一标识
    sentenceId   string 

}

func NewTaobaoAilabAicloudTopFeedlistDeleteRequest() *TaobaoAilabAicloudTopFeedlistDeleteRequest{
    return &TaobaoAilabAicloudTopFeedlistDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.feedlist.delete"
}

func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetSentenceId(sentenceId string) error {
    r.sentenceId = sentenceId
    r.Set("sentence_id", sentenceId)
    return nil
}

func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetSentenceId() string {
    return r.sentenceId
}

