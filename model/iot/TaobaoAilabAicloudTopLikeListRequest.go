package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
列出收藏列表 APIRequest
taobao.ailab.aicloud.top.like.list

列出收藏列表
*/
type TaobaoAilabAicloudTopLikeListRequest struct {
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
    param1   string 

    // 页码 从0起
    param2   int64 

    // 每页条目个数
    param3   int64 

}

func NewTaobaoAilabAicloudTopLikeListRequest() *TaobaoAilabAicloudTopLikeListRequest{
    return &TaobaoAilabAicloudTopLikeListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.list"
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopLikeListRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopLikeListRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopLikeListRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopLikeListRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopLikeListRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopLikeListRequest) SetParam2(param2 int64) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetParam2() int64 {
    return r.param2
}

func (r *TaobaoAilabAicloudTopLikeListRequest) SetParam3(param3 int64) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

func (r TaobaoAilabAicloudTopLikeListRequest) GetParam3() int64 {
    return r.param3
}

