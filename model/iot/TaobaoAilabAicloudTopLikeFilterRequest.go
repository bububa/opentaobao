package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
过滤列表歌曲存在于收藏列表的 API请求
taobao.ailab.aicloud.top.like.filter

过滤出传入列表歌曲存在于收藏列表的
*/
type TaobaoAilabAicloudTopLikeFilterRequest struct {
    model.Params
    // 账户体系隔离
    _schema   string
    // 用户ID，此处传入第三方账户体系的用户id
    _userId   string
    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    _utdId   string
    // 扩展信息，用于存放APP类型等
    _ext   string
    // 音频收藏类型, 四种类型：music,children_song,program,story
    _type   string
    // 传入的歌曲列表
    _mediaItems   []MediaItem
}

// 初始化TaobaoAilabAicloudTopLikeFilterRequest对象
func NewTaobaoAilabAicloudTopLikeFilterRequest() *TaobaoAilabAicloudTopLikeFilterRequest{
    return &TaobaoAilabAicloudTopLikeFilterRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.filter"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetExt() string {
    return r._ext
}
// Type Setter
// 音频收藏类型, 四种类型：music,children_song,program,story
func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetType() string {
    return r._type
}
// MediaItems Setter
// 传入的歌曲列表
func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetMediaItems(_mediaItems []MediaItem) error {
    r._mediaItems = _mediaItems
    r.Set("media_items", _mediaItems)
    return nil
}

// MediaItems Getter
func (r TaobaoAilabAicloudTopLikeFilterRequest) GetMediaItems() []MediaItem {
    return r._mediaItems
}
