package xhotelofficial

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住取消结账 APIRequest
taobao.xhotel.order.official.settle.cancel

用于官网信用住取消结账
*/
type TaobaoXhotelOrderOfficialSettleCancelRequest struct {
    model.Params

    // 阿里旅行订单号，淘宝订单号或外部订单号二选一必填
    tid   int64 

    // 取消结账的原因
    reason   string 

    // 外部订单号，和tid二选一必填
    outId   string 

    // 暂时无意义，无需传入
    notifyUrl   string 

    // 请求流水号
    outUuid   string 

}

func NewTaobaoXhotelOrderOfficialSettleCancelRequest() *TaobaoXhotelOrderOfficialSettleCancelRequest{
    return &TaobaoXhotelOrderOfficialSettleCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderOfficialSettleCancelRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.settle.cancel"
}

func (r TaobaoXhotelOrderOfficialSettleCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderOfficialSettleCancelRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelOrderOfficialSettleCancelRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelOrderOfficialSettleCancelRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r TaobaoXhotelOrderOfficialSettleCancelRequest) GetReason() string {
    return r.reason
}

func (r *TaobaoXhotelOrderOfficialSettleCancelRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoXhotelOrderOfficialSettleCancelRequest) GetOutId() string {
    return r.outId
}

func (r *TaobaoXhotelOrderOfficialSettleCancelRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

func (r TaobaoXhotelOrderOfficialSettleCancelRequest) GetNotifyUrl() string {
    return r.notifyUrl
}

func (r *TaobaoXhotelOrderOfficialSettleCancelRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

func (r TaobaoXhotelOrderOfficialSettleCancelRequest) GetOutUuid() string {
    return r.outUuid
}

