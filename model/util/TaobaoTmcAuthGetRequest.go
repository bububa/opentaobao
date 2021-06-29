package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
TMC授权token API请求
taobao.tmc.auth.get

TMC连接授权Token
*/
type TaobaoTmcAuthGetRequest struct {
    model.Params
    // tmc组名
    _group   string
}

// 初始化TaobaoTmcAuthGetRequest对象
func NewTaobaoTmcAuthGetRequest() *TaobaoTmcAuthGetRequest{
    return &TaobaoTmcAuthGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcAuthGetRequest) GetApiMethodName() string {
    return "taobao.tmc.auth.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcAuthGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Group Setter
// tmc组名
func (r *TaobaoTmcAuthGetRequest) SetGroup(_group string) error {
    r._group = _group
    r.Set("group", _group)
    return nil
}

// Group Getter
func (r TaobaoTmcAuthGetRequest) GetGroup() string {
    return r._group
}
