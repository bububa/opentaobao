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
    _distributor   string
    // 淘宝订单号
    _tid   int64
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
func (r *TaobaoOpenmallTradeConfirmRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeConfirmRequest) GetDistributor() string {
    return r._distributor
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeConfirmRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeConfirmRequest) GetTid() int64 {
    return r._tid
}
