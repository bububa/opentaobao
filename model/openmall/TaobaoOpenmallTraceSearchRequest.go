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
type TaobaoOpenmallTraceSearchRequest struct {
    model.Params
    // 签约支付宝代扣的账号
    distributor   string
    // 淘宝订单编号
    tid   int64
}

// 初始化TaobaoOpenmallTraceSearchRequest对象
func NewTaobaoOpenmallTraceSearchRequest() *TaobaoOpenmallTraceSearchRequest{
    return &TaobaoOpenmallTraceSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTraceSearchRequest) GetApiMethodName() string {
    return "taobao.openmall.trace.search"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTraceSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 签约支付宝代扣的账号
func (r *TaobaoOpenmallTraceSearchRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTraceSearchRequest) GetDistributor() string {
    return r.distributor
}
// Tid Setter
// 淘宝订单编号
func (r *TaobaoOpenmallTraceSearchRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTraceSearchRequest) GetTid() int64 {
    return r.tid
}
