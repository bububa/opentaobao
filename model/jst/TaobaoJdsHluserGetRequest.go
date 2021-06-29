package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单全链路用户信息获取 API请求
taobao.jds.hluser.get

订单全链路用户信息获取
*/
type TaobaoJdsHluserGetRequest struct {
    model.Params
}

// 初始化TaobaoJdsHluserGetRequest对象
func NewTaobaoJdsHluserGetRequest() *TaobaoJdsHluserGetRequest{
    return &TaobaoJdsHluserGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsHluserGetRequest) GetApiMethodName() string {
    return "taobao.jds.hluser.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsHluserGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
