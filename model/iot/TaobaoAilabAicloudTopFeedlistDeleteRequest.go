package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除单条对话流信息 API请求
taobao.ailab.aicloud.top.feedlist.delete

删除指定的某一条对话流信息
*/
type TaobaoAilabAicloudTopFeedlistDeleteRequest struct {
    model.Params
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 消息的唯一标识
    _sentenceId   string
}

// 初始化TaobaoAilabAicloudTopFeedlistDeleteRequest对象
func NewTaobaoAilabAicloudTopFeedlistDeleteRequest() *TaobaoAilabAicloudTopFeedlistDeleteRequest{
    return &TaobaoAilabAicloudTopFeedlistDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.feedlist.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetExt() string {
    return r._ext
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetUtdId() string {
    return r._utdId
}
// SentenceId Setter
// 消息的唯一标识
func (r *TaobaoAilabAicloudTopFeedlistDeleteRequest) SetSentenceId(_sentenceId string) error {
    r._sentenceId = _sentenceId
    r.Set("sentence_id", _sentenceId)
    return nil
}

// SentenceId Getter
func (r TaobaoAilabAicloudTopFeedlistDeleteRequest) GetSentenceId() string {
    return r._sentenceId
}
