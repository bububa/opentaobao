package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单取消结算接口 APIRequest
taobao.xhotel.order.alipayface.cancelsettle

信用住订单由于客人为出现等原因，最终取消结算,一定要在结算后2个小时之内调用，否则不会成功。
*/
type TaobaoXhotelOrderAlipayfaceCancelsettleRequest struct {
    model.Params

    // 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
    tid   int64 

    // 取消结账的原因
    reason   string 

    // 外部订单号，和tid二选一必填
    outId   string 

}

func NewTaobaoXhotelOrderAlipayfaceCancelsettleRequest() *TaobaoXhotelOrderAlipayfaceCancelsettleRequest{
    return &TaobaoXhotelOrderAlipayfaceCancelsettleRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.cancelsettle"
}

func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderAlipayfaceCancelsettleRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelOrderAlipayfaceCancelsettleRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetReason() string {
    return r.reason
}

func (r *TaobaoXhotelOrderAlipayfaceCancelsettleRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCancelsettleRequest) GetOutId() string {
    return r.outId
}

