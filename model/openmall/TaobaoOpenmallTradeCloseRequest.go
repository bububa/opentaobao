package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭订单 API请求
taobao.openmall.trade.close

关闭订单
*/
type TaobaoOpenmallTradeCloseRequest struct {
    model.Params
    // 分销者信息
    _distributor   string
    // 关单原因
    _reason   string
    // 淘宝订单号
    _tid   int64
}

// 初始化TaobaoOpenmallTradeCloseRequest对象
func NewTaobaoOpenmallTradeCloseRequest() *TaobaoOpenmallTradeCloseRequest{
    return &TaobaoOpenmallTradeCloseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeCloseRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.close"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeCloseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 分销者信息
func (r *TaobaoOpenmallTradeCloseRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeCloseRequest) GetDistributor() string {
    return r._distributor
}
// Reason Setter
// 关单原因
func (r *TaobaoOpenmallTradeCloseRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r TaobaoOpenmallTradeCloseRequest) GetReason() string {
    return r._reason
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeCloseRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeCloseRequest) GetTid() int64 {
    return r._tid
}
