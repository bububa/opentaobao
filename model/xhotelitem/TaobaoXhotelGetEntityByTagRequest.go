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
type TaobaoXhotelGetEntityByTagRequest struct {
    model.Params
    // 标签
    tag   int64
    // 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
    tokenStr   string
}

// 初始化TaobaoXhotelGetEntityByTagRequest对象
func NewTaobaoXhotelGetEntityByTagRequest() *TaobaoXhotelGetEntityByTagRequest{
    return &TaobaoXhotelGetEntityByTagRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelGetEntityByTagRequest) GetApiMethodName() string {
    return "taobao.xhotel.get.entity.by.tag"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelGetEntityByTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tag Setter
// 标签
func (r *TaobaoXhotelGetEntityByTagRequest) SetTag(tag int64) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

// Tag Getter
func (r TaobaoXhotelGetEntityByTagRequest) GetTag() int64 {
    return r.tag
}
// TokenStr Setter
// 查询token，填入上一页查询的返回结果，只能按顺序单线程查询
func (r *TaobaoXhotelGetEntityByTagRequest) SetTokenStr(tokenStr string) error {
    r.tokenStr = tokenStr
    r.Set("token_str", tokenStr)
    return nil
}

// TokenStr Getter
func (r TaobaoXhotelGetEntityByTagRequest) GetTokenStr() string {
    return r.tokenStr
}
