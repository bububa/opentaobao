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
    _distributor   string
    // 淘宝订单号
    _tid   int64
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
func (r *TaobaoOpenmallTradeGetRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeGetRequest) GetDistributor() string {
    return r._distributor
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeGetRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeGetRequest) GetTid() int64 {
    return r._tid
}
