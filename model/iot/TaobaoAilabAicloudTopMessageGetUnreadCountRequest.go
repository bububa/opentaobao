package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取未读的消息数量 APIRequest
taobao.ailab.aicloud.top.message.get.unread.count

开放获取未读留言数量的接口
*/
type TaobaoAilabAicloudTopMessageGetUnreadCountRequest struct {
    model.Params

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

}

func NewTaobaoAilabAicloudTopMessageGetUnreadCountRequest() *TaobaoAilabAicloudTopMessageGetUnreadCountRequest{
    return &TaobaoAilabAicloudTopMessageGetUnreadCountRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.get.unread.count"
}

func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetExt() string {
    return r.ext
}

