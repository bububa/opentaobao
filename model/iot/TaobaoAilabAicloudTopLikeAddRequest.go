package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加收藏 API请求
taobao.ailab.aicloud.top.like.add

将制定内容加入收藏
*/
type TaobaoAilabAicloudTopLikeAddRequest struct {
    model.Params
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 账户体系隔离
    _schema   string
    // 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
    _type   string
    // 来源
    _source   string
    // 收藏的资源的ID
    _itemId   string
    // 内容，必须要是一个json格式：{"song":"走过1999","singer":"张学友","album":"走过1999"}
    _content   string
}

// 初始化TaobaoAilabAicloudTopLikeAddRequest对象
func NewTaobaoAilabAicloudTopLikeAddRequest() *TaobaoAilabAicloudTopLikeAddRequest{
    return &TaobaoAilabAicloudTopLikeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopLikeAddRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopLikeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetExt() string {
    return r._ext
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetUtdId() string {
    return r._utdId
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetUserId() string {
    return r._userId
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetSchema() string {
    return r._schema
}
// Type Setter
// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetType() string {
    return r._type
}
// Source Setter
// 来源
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetSource(_source string) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetSource() string {
    return r._source
}
// ItemId Setter
// 收藏的资源的ID
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetItemId(_itemId string) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetItemId() string {
    return r._itemId
}
// Content Setter
// 内容，必须要是一个json格式：{"song":"走过1999","singer":"张学友","album":"走过1999"}
func (r *TaobaoAilabAicloudTopLikeAddRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoAilabAicloudTopLikeAddRequest) GetContent() string {
    return r._content
}
