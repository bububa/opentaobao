package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过订单获取对应买家的openUID API请求
taobao.openuid.get.bytrade

通过订单获取对应买家的openUID,需要卖家授权
*/
type TaobaoOpenuidGetBytradeRequest struct {
    model.Params
    // 订单ID
    _tid   int64
}

// 初始化TaobaoOpenuidGetBytradeRequest对象
func NewTaobaoOpenuidGetBytradeRequest() *TaobaoOpenuidGetBytradeRequest{
    return &TaobaoOpenuidGetBytradeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenuidGetBytradeRequest) GetApiMethodName() string {
    return "taobao.openuid.get.bytrade"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenuidGetBytradeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 订单ID
func (r *TaobaoOpenuidGetBytradeRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenuidGetBytradeRequest) GetTid() int64 {
    return r._tid
}
