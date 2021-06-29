package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认收货 API请求
taobao.openmall.trade.confirm

确认订单收货
*/
type TaobaoOpenmallTradeConfirmRequest struct {
    model.Params
    // 分销者信息
    distributor   string
    // 淘宝订单号
    tid   int64
}

// 初始化TaobaoOpenmallTradeConfirmRequest对象
func NewTaobaoOpenmallTradeConfirmRequest() *TaobaoOpenmallTradeConfirmRequest{
    return &TaobaoOpenmallTradeConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeConfirmRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 分销者信息
func (r *TaobaoOpenmallTradeConfirmRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeConfirmRequest) GetDistributor() string {
    return r.distributor
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeConfirmRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeConfirmRequest) GetTid() int64 {
    return r.tid
}
