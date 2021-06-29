package xhotelofficial

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
官网信用住订单取消 API请求
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

// 初始化TaobaoXhotelOrderOfficialCancelRequest对象
func NewTaobaoXhotelOrderOfficialCancelRequest() *TaobaoXhotelOrderOfficialCancelRequest{
    return &TaobaoXhotelOrderOfficialCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderOfficialCancelRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.official.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderOfficialCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单号,必选
func (r *TaobaoXhotelOrderOfficialCancelRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderOfficialCancelRequest) GetTid() int64 {
    return r.tid
}
// ReasonText Setter
// 原因描述
func (r *TaobaoXhotelOrderOfficialCancelRequest) SetReasonText(reasonText string) error {
    r.reasonText = reasonText
    r.Set("reason_text", reasonText)
    return nil
}

// ReasonText Getter
func (r TaobaoXhotelOrderOfficialCancelRequest) GetReasonText() string {
    return r.reasonText
}
// OutId Setter
// 外部订单号
func (r *TaobaoXhotelOrderOfficialCancelRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderOfficialCancelRequest) GetOutId() string {
    return r.outId
}
// OutUuid Setter
// 请求流水号（必须传入）
func (r *TaobaoXhotelOrderOfficialCancelRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderOfficialCancelRequest) GetOutUuid() string {
    return r.outUuid
}
// NotifyUrl Setter
// 暂无意义，无需传入
func (r *TaobaoXhotelOrderOfficialCancelRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

// NotifyUrl Getter
func (r TaobaoXhotelOrderOfficialCancelRequest) GetNotifyUrl() string {
    return r.notifyUrl
}
