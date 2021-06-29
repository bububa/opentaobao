package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下信用住订单取消 API请求
taobao.xhotel.order.alipayface.cancel

提供给卖家进行线下信用住的订单取消。此接口仅仅支持线下信用住订单的取消
*/
type TaobaoXhotelOrderAlipayfaceCancelRequest struct {
    model.Params
    // 淘宝订单ID，必选
    tid   int64
    // 原因描述
    reasonText   string
    // 外部订单号
    outId   string
    // 预留后续用
    notifyUrl   string
    // 请求流水号
    outUuid   string
}

// 初始化TaobaoXhotelOrderAlipayfaceCancelRequest对象
func NewTaobaoXhotelOrderAlipayfaceCancelRequest() *TaobaoXhotelOrderAlipayfaceCancelRequest{
    return &TaobaoXhotelOrderAlipayfaceCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCancelRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝订单ID，必选
func (r *TaobaoXhotelOrderAlipayfaceCancelRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderAlipayfaceCancelRequest) GetTid() int64 {
    return r.tid
}
// ReasonText Setter
// 原因描述
func (r *TaobaoXhotelOrderAlipayfaceCancelRequest) SetReasonText(reasonText string) error {
    r.reasonText = reasonText
    r.Set("reason_text", reasonText)
    return nil
}

// ReasonText Getter
func (r TaobaoXhotelOrderAlipayfaceCancelRequest) GetReasonText() string {
    return r.reasonText
}
// OutId Setter
// 外部订单号
func (r *TaobaoXhotelOrderAlipayfaceCancelRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

// OutId Getter
func (r TaobaoXhotelOrderAlipayfaceCancelRequest) GetOutId() string {
    return r.outId
}
// NotifyUrl Setter
// 预留后续用
func (r *TaobaoXhotelOrderAlipayfaceCancelRequest) SetNotifyUrl(notifyUrl string) error {
    r.notifyUrl = notifyUrl
    r.Set("notify_url", notifyUrl)
    return nil
}

// NotifyUrl Getter
func (r TaobaoXhotelOrderAlipayfaceCancelRequest) GetNotifyUrl() string {
    return r.notifyUrl
}
// OutUuid Setter
// 请求流水号
func (r *TaobaoXhotelOrderAlipayfaceCancelRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderAlipayfaceCancelRequest) GetOutUuid() string {
    return r.outUuid
}
