package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分查询接口 APIRequest
taobao.jst.interactive.point.query

查询用户的互动积分
*/
type TaobaoJstInteractivePointQueryRequest struct {
    model.Params

}

func NewTaobaoJstInteractivePointQueryRequest() *TaobaoJstInteractivePointQueryRequest{
    return &TaobaoJstInteractivePointQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractivePointQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.point.query"
}

func (r TaobaoJstInteractivePointQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


