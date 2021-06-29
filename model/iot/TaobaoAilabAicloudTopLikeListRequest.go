package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
列出收藏列表 API请求
taobao.ailab.aicloud.top.like.list

列出收藏列表
*/
type TaobaoAilabAicloudTopLikeListRequest struct {
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
    _param1   string
    // 页码 从0起
    _param2   int64
    // 每页条目个数
    _param3   int64
}

// 初始化TaobaoAilabAicloudTopLikeListRequest对象
func NewTaobaoAilabAicloudTopLikeListRequest() *TaobaoAilabAicloudTopLikeListRequest{
    return &TaobaoAilabAicloudTopLikeListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopLikeListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.like.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopLikeListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopLikeListRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopLikeListRequest) GetExt() string {
    return r._ext
}
// UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopLikeListRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopLikeListRequest) GetUtdId() string {
    return r._utdId
}
// UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopLikeListRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopLikeListRequest) GetUserId() string {
    return r._userId
}
// Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopLikeListRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopLikeListRequest) GetSchema() string {
    return r._schema
}
// Param1 Setter
// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
func (r *TaobaoAilabAicloudTopLikeListRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopLikeListRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// 页码 从0起
func (r *TaobaoAilabAicloudTopLikeListRequest) SetParam2(_param2 int64) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopLikeListRequest) GetParam2() int64 {
    return r._param2
}
// Param3 Setter
// 每页条目个数
func (r *TaobaoAilabAicloudTopLikeListRequest) SetParam3(_param3 int64) error {
    r._param3 = _param3
    r.Set("param3", _param3)
    return nil
}

// Param3 Getter
func (r TaobaoAilabAicloudTopLikeListRequest) GetParam3() int64 {
    return r._param3
}
