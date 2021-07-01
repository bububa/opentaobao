package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取Openmall订单物流流转信息 API请求
taobao.openmall.trace.search

获取Openmall订单物流流转信息
*/
type TaobaoOpenmallTraceSearchAPIRequest struct {
    model.Params
    // 签约支付宝代扣的账号
    _distributor   string
    // 淘宝订单编号
    _tid   int64
}

// 初始化TaobaoOpenmallTraceSearchAPIRequest对象
func NewTaobaoOpenmallTraceSearchRequest() *TaobaoOpenmallTraceSearchAPIRequest{
    return &TaobaoOpenmallTraceSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTraceSearchAPIRequest) GetApiMethodName() string {
    return "taobao.openmall.trace.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTraceSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 签约支付宝代扣的账号
func (r *TaobaoOpenmallTraceSearchAPIRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTraceSearchAPIRequest) GetDistributor() string {
    return r._distributor
}
// Tid Setter
// 淘宝订单编号
func (r *TaobaoOpenmallTraceSearchAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTraceSearchAPIRequest) GetTid() int64 {
    return r._tid
}
