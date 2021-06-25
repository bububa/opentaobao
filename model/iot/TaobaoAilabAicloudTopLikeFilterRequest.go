package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
过滤列表歌曲存在于收藏列表的 APIRequest
taobao.ailab.aicloud.top.like.filter

过滤出传入列表歌曲存在于收藏列表的
*/
type TaobaoAilabAicloudTopLikeFilterRequest struct {
    model.Params

    // 账户体系隔离
    schema   string 

    // 用户ID，此处传入第三方账户体系的用户id
    userId   string 

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    utdId   string 

    // 扩展信息，用于存放APP类型等
    ext   string 

    // 音频收藏类型, 四种类型：music,children_song,program,story
    type   string 

    // 传入的歌曲列表
    mediaItems   []MediaItem 

}

func NewTaobaoAilabAicloudTopLikeFilterRequest() *TaobaoAilabAicloudTopLikeFilterRequest{
    return &TaobaoAilabAicloudTopLikeFilterRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.filter"
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetType() string {
    return r.type
}

func (r *TaobaoAilabAicloudTopLikeFilterRequest) SetMediaItems(mediaItems []MediaItem) error {
    r.mediaItems = mediaItems
    r.Set("media_items", mediaItems)
    return nil
}

func (r TaobaoAilabAicloudTopLikeFilterRequest) GetMediaItems() []MediaItem {
    return r.mediaItems
}

