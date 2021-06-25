package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加收藏 APIRequest
taobao.ailab.aicloud.top.like.add

将制定内容加入收藏
*/
type TaobaoAilabAicloudTopLikeAddRequest struct {
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

    // 收藏的资源的ID
    itemId   string 

    // 内容，必须要是一个json格式：{"song":"走过1999","singer":"张学友","album":"走过1999"}
    content   string 

}

func NewTaobaoAilabAicloudTopLikeAddRequest() *TaobaoAilabAicloudTopLikeAddRequest{
    return &TaobaoAilabAicloudTopLikeAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.add"
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopLikeAddRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopLikeAddRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopLikeAddRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopLikeAddRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopLikeAddRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetType() string {
    return r.type
}

func (r *TaobaoAilabAicloudTopLikeAddRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetSource() string {
    return r.source
}

func (r *TaobaoAilabAicloudTopLikeAddRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetItemId() string {
    return r.itemId
}

func (r *TaobaoAilabAicloudTopLikeAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoAilabAicloudTopLikeAddRequest) GetContent() string {
    return r.content
}

