package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单取消结算接口 API请求
taobao.xhotel.order.alipayface.cancelsettle

信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。
*/
type TaobaoXhotelOrderAlipayfaceCancelsettleRequest struct {
    model.Params
    // 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
    _tid   int64
    // 取消结账的原因
    _reason   string
    // 外部订单号，和tid二选一必填
    _outId   string
}

// 初始化TaobaoXhotelOrderAlipayfaceCancelsettleRequest对象
func NewTaobaoXhotelOrderAlipayfaceCancelsettleRequest() *TaobaoXhotelOrderAlipayfaceCancelsettleRequest{
    return &TaobaoXhotelOrderAlipayfaceCancelsettleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.cancelsettle"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
func (r *TaobaoXhotelOrderAlipayfaceCancelsettleRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetTid() int64 {
    return r._tid
}
// Reason Setter
// 取消结账的原因
func (r *TaobaoXhotelOrderAlipayfaceCancelsettleRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetReason() string {
    return r._reason
}
// OutId Setter
// 外部订单号，和tid二选一必填
func (r *TaobaoXhotelOrderAlipayfaceCancelsettleRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetOutId() string {
    return r._outId
}
