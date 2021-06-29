package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单条订单跟踪详情 API请求
taobao.jds.trade.traces.get

获取聚石塔数据共享的交易全链路信息
*/
type TaobaoJdsTradeTracesGetRequest struct {
    model.Params
    // 淘宝的订单tid
    _tid   int64
}

// 初始化TaobaoJdsTradeTracesGetRequest对象
func NewTaobaoJdsTradeTracesGetRequest() *TaobaoJdsTradeTracesGetRequest{
    return &TaobaoJdsTradeTracesGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJdsTradeTracesGetRequest) GetApiMethodName() string {
    return "taobao.jds.trade.traces.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJdsTradeTracesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝的订单tid
func (r *TaobaoJdsTradeTracesGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoJdsTradeTracesGetRequest) GetTid() int64 {
    return r._tid
}
