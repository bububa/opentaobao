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
type TaobaoJdsHluserGetAPIRequest struct {
    model.Params
}

// 初始化TaobaoJdsHluserGetAPIRequest对象
func NewTaobaoJdsHluserGetRequest() *TaobaoJdsHluserGetAPIRequest{
    return &TaobaoJdsHluserGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsHluserGetAPIRequest) GetApiMethodName() string {
    return "taobao.jds.hluser.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsHluserGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
