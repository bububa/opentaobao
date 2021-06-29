package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分查询接口 API请求
taobao.jst.interactive.point.query

查询用户的互动积分
*/
type TaobaoJstInteractivePointQueryRequest struct {
    model.Params
}

// 初始化TaobaoJstInteractivePointQueryRequest对象
func NewTaobaoJstInteractivePointQueryRequest() *TaobaoJstInteractivePointQueryRequest{
    return &TaobaoJstInteractivePointQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractivePointQueryRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.point.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractivePointQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
