package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消收藏 APIRequest
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

func NewTaobaoAilabAicloudTopLikeDeleteRequest() *TaobaoAilabAicloudTopLikeDeleteRequest{
    return &TaobaoAilabAicloudTopLikeDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.delete"
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetType() string {
    return r.type
}

func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetSource() string {
    return r.source
}

func (r *TaobaoAilabAicloudTopLikeDeleteRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeDeleteRequest) GetItemId() string {
    return r.itemId
}

