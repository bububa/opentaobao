package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消收藏 API请求
taobao.ailab.aicloud.top.like.delete

取消收藏
*/
type TaobaoAilabAicloudTopLikeDeleteRequest struct {
    model.Params
    // 扩展信息，用于存放APP类型等
    ext   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string
    // 用户ID，此处传入第三方账户体系的用户id
    userId   string
    // 账户体系隔离
    schema   string
    // 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
    type   string
    // 来源
    source   string
    // 资源的记录ID
    itemId   string
}

// 初始化TaobaoAilabAicloudTopLikeDeleteRequest对象
func NewTaobaoAilabAicloudTopLikeDeleteRequest() *TaobaoAilabAicloudTopLikeDeleteRequest{
    return &TaobaoAilabAicloudTopLikeDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetExt() string {
    return r.ext
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetUtdId() string {
    return r.utdId
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetUserId() string {
    return r.userId
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetSchema() string {
    return r.schema
}
// Type Setter
// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetType() string {
    return r.type
}
// Source Setter
// 来源
func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetSource() string {
    return r.source
}
// ItemId Setter
// 资源的记录ID
func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetItemId() string {
    return r.itemId
}
