package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店订单发货接口 API请求
taobao.xhotel.order.update

卖家确认订单或者取消订单，适用于预付、面付、信用住订单
*/
type TaobaoXhotelOrderUpdateRequest struct {
    model.Params
    // 订单号
    _tid   int64
    // 操作的类型：1.确认无房（取消预订，710发送短信提醒买家申请退款）2.确认预订 3.入住 4.离店 5.noshow 6.关单
    _optType   int64
    // 是否把代理直签的订单同步到酒店，Y为同步，N不同步
    _syncToHotel   string
    // 退款费用
    _refundFee   int64
    // 取消类型，6 代表的是用户取消，reasonType=7代表的是小二协商
    _reasonType   int64
    // 开票金额
    _invoiceAmount   int64
}

// 初始化TaobaoXhotelOrderUpdateRequest对象
func NewTaobaoXhotelOrderUpdateRequest() *TaobaoXhotelOrderUpdateRequest{
    return &TaobaoXhotelOrderUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 订单号
func (r *TaobaoXhotelOrderUpdateRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoXhotelOrderUpdateRequest) GetTid() int64 {
    return r._tid
}
// OptType Setter
// 操作的类型：1.确认无房（取消预订，710发送短信提醒买家申请退款）2.确认预订 3.入住 4.离店 5.noshow 6.关单
func (r *TaobaoXhotelOrderUpdateRequest) SetOptType(_optType int64) error {
    r._optType = _optType
    r.Set("opt_type", _optType)
    return nil
}

// OptType Getter
func (r TaobaoXhotelOrderUpdateRequest) GetOptType() int64 {
    return r._optType
}
// SyncToHotel Setter
// 是否把代理直签的订单同步到酒店，Y为同步，N不同步
func (r *TaobaoXhotelOrderUpdateRequest) SetSyncToHotel(_syncToHotel string) error {
    r._syncToHotel = _syncToHotel
    r.Set("sync_to_hotel", _syncToHotel)
    return nil
}

// SyncToHotel Getter
func (r TaobaoXhotelOrderUpdateRequest) GetSyncToHotel() string {
    return r._syncToHotel
}
// RefundFee Setter
// 退款费用
func (r *TaobaoXhotelOrderUpdateRequest) SetRefundFee(_refundFee int64) error {
    r._refundFee = _refundFee
    r.Set("refund_fee", _refundFee)
    return nil
}

// RefundFee Getter
func (r TaobaoXhotelOrderUpdateRequest) GetRefundFee() int64 {
    return r._refundFee
}
// ReasonType Setter
// 取消类型，6 代表的是用户取消，reasonType=7代表的是小二协商
func (r *TaobaoXhotelOrderUpdateRequest) SetReasonType(_reasonType int64) error {
    r._reasonType = _reasonType
    r.Set("reason_type", _reasonType)
    return nil
}

// ReasonType Getter
func (r TaobaoXhotelOrderUpdateRequest) GetReasonType() int64 {
    return r._reasonType
}
// InvoiceAmount Setter
// 开票金额
func (r *TaobaoXhotelOrderUpdateRequest) SetInvoiceAmount(_invoiceAmount int64) error {
    r._invoiceAmount = _invoiceAmount
    r.Set("invoice_amount", _invoiceAmount)
    return nil
}

// InvoiceAmount Getter
func (r TaobaoXhotelOrderUpdateRequest) GetInvoiceAmount() int64 {
    return r._invoiceAmount
}
