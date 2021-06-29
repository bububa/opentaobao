package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取未读的消息数量 API请求
taobao.ailab.aicloud.top.message.get.unread.count

开放获取未读留言数量的接口
*/
type TaobaoAilabAicloudTopMessageGetUnreadCountRequest struct {
    model.Params
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
}

// 初始化TaobaoAilabAicloudTopMessageGetUnreadCountRequest对象
func NewTaobaoAilabAicloudTopMessageGetUnreadCountRequest() *TaobaoAilabAicloudTopMessageGetUnreadCountRequest{
    return &TaobaoAilabAicloudTopMessageGetUnreadCountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.get.unread.count"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountRequest) GetExt() string {
    return r._ext
}
