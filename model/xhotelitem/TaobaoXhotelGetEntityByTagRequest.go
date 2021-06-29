package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据标签查询实体 APIRequest
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

func NewTaobaoXhotelGetEntityByTagRequest() *TaobaoXhotelGetEntityByTagRequest{
    return &TaobaoXhotelGetEntityByTagRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelGetEntityByTagRequest) GetApiMethodName() string {
    return "taobao.xhotel.get.entity.by.tag"
}

func (r TaobaoXhotelGetEntityByTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelGetEntityByTagRequest) SetTag(tag int64) error {
    r.tag = tag
    r.Set("tag", tag)
    return nil
}

func (r TaobaoXhotelGetEntityByTagRequest) GetTag() int64 {
    return r.tag
}

func (r *TaobaoXhotelGetEntityByTagRequest) SetTokenStr(tokenStr string) error {
    r.tokenStr = tokenStr
    r.Set("token_str", tokenStr)
    return nil
}

func (r TaobaoXhotelGetEntityByTagRequest) GetTokenStr() string {
    return r.tokenStr
}

