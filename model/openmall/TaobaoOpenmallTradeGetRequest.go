package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单详情 API请求
taobao.openmall.trade.get

查询订单详情
*/
type TaobaoOpenmallTradeGetRequest struct {
    model.Params
    // 分销者信息
    distributor   string
    // 淘宝订单号
    tid   int64
}

// 初始化TaobaoOpenmallTradeGetRequest对象
func NewTaobaoOpenmallTradeGetRequest() *TaobaoOpenmallTradeGetRequest{
    return &TaobaoOpenmallTradeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeGetRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 分销者信息
func (r *TaobaoOpenmallTradeGetRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeGetRequest) GetDistributor() string {
    return r.distributor
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeGetRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeGetRequest) GetTid() int64 {
    return r.tid
}
