package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据标签查询实体 API请求
taobao.xhotel.get.entity.by.tag

根据标签查询实体
*/
type TaobaoXhotelGetEntityByTagAPIRequest struct {
    model.Params
    // 标签
    _tag   int64
    // 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
    _tokenStr   string
}

// 初始化TaobaoXhotelGetEntityByTagAPIRequest对象
func NewTaobaoXhotelGetEntityByTagRequest() *TaobaoXhotelGetEntityByTagAPIRequest{
    return &TaobaoXhotelGetEntityByTagAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.get.entity.by.tag"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tag Setter
// 标签
func (r *TaobaoXhotelGetEntityByTagAPIRequest) SetTag(_tag int64) error {
    r._tag = _tag
    r.Set("tag", _tag)
    return nil
}

// Tag Getter
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetTag() int64 {
    return r._tag
}
// TokenStr Setter
// 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
func (r *TaobaoXhotelGetEntityByTagAPIRequest) SetTokenStr(_tokenStr string) error {
    r._tokenStr = _tokenStr
    r.Set("token_str", _tokenStr)
    return nil
}

// TokenStr Getter
func (r TaobaoXhotelGetEntityByTagAPIRequest) GetTokenStr() string {
    return r._tokenStr
}
