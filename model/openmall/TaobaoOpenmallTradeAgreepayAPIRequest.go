package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openmall订单支付 API请求
taobao.openmall.trade.agreepay

openmall订单支付
*/
type TaobaoOpenmallTradeAgreepayAPIRequest struct {
    model.Params
    // 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
    _distributor   string
    // 淘宝交易单号
    _tid   int64
}

// 初始化TaobaoOpenmallTradeAgreepayAPIRequest对象
func NewTaobaoOpenmallTradeAgreepayRequest() *TaobaoOpenmallTradeAgreepayAPIRequest{
    return &TaobaoOpenmallTradeAgreepayAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.agreepay"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
func (r *TaobaoOpenmallTradeAgreepayAPIRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetDistributor() string {
    return r._distributor
}
// Tid Setter
// 淘宝交易单号
func (r *TaobaoOpenmallTradeAgreepayAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetTid() int64 {
    return r._tid
}
