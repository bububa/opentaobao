package xhotelofficial

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住订单取消 APIRequest
taobao.xhotel.order.official.cancel

官网信用住订单取消
*/
type TaobaoXhotelOrderOfficialCancelRequest struct {
    model.Params

    // 淘宝订单号,必选
    tid   int64 

    // 原因描述
    reasonText   string 

    // 外部订单号
    outId   string 

    // 请求流水号（必须传入）
    outUuid   string 

    // 暂无意义，无需传入
    notifyUrl   string 

}

func NewTaobaoXhotelOrderOfficialCancelRequest() *TaobaoXhotelOrderOfficialCancelRequest{
    return &TaobaoXhotelOrderOfficialCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderOfficialCancelRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.cancel"
}

func (r TaobaoXhotelOrderOfficialCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderOfficialCancelRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelOrderOfficialCancelRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelOrderOfficialCancelRequest) SetReasonText(reasonText string) error {
    r.reasonText = reasonText
    r.Set("reason_text", reasonText)
    return nil
}

func (r TaobaoXhotelOrderOfficialCancelRequest) GetReasonText() string {
    return r.reasonText
}

func (r *TaobaoXhotelOrderOfficialCancelRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoXhotelOrderOfficialCancelRequest) GetOutId() string {
    return r.outId
}

func (r *TaobaoXhotelOrderOfficialCancelRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

func (r TaobaoXhotelOrderOfficialCancelRequest) GetOutUuid() string {
    return r.outUuid
}

func (r *TaobaoXhotelOrderOfficialCancelRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

func (r TaobaoXhotelOrderOfficialCancelRequest) GetNotifyUrl() string {
    return r.notifyUrl
}

