package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店订单发货接口 APIRequest
taobao.xhotel.order.update

卖家确认订单或者取消订单，适用于预付、面付、信用住订单
*/
type TaobaoXhotelOrderUpdateRequest struct {
    model.Params

    // 订单号
    tid   int64 

    // 操作的类型：1.确认无房（取消预订，710发送短信提醒买家申请退款）2.确认预订 3.入住 4.离店 5.noshow 6.关单
    optType   int64 

    // 是否把代理直签的订单同步到酒店，Y为同步，N不同步
    syncToHotel   string 

    // 退款费用
    refundFee   int64 

    // 取消类型，6 代表的是用户取消，reasonType=7代表的是小二协商
    reasonType   int64 

    // 开票金额
    invoiceAmount   int64 

}

func NewTaobaoXhotelOrderUpdateRequest() *TaobaoXhotelOrderUpdateRequest{
    return &TaobaoXhotelOrderUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.update"
}

func (r TaobaoXhotelOrderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoXhotelOrderUpdateRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoXhotelOrderUpdateRequest) SetOptType(optType int64) error {
    r.optType = optType
    r.Set("opt_type", optType)
    return nil
}

func (r TaobaoXhotelOrderUpdateRequest) GetOptType() int64 {
    return r.optType
}

func (r *TaobaoXhotelOrderUpdateRequest) SetSyncToHotel(syncToHotel string) error {
    r.syncToHotel = syncToHotel
    r.Set("sync_to_hotel", syncToHotel)
    return nil
}

func (r TaobaoXhotelOrderUpdateRequest) GetSyncToHotel() string {
    return r.syncToHotel
}

func (r *TaobaoXhotelOrderUpdateRequest) SetRefundFee(refundFee int64) error {
    r.refundFee = refundFee
    r.Set("refund_fee", refundFee)
    return nil
}

func (r TaobaoXhotelOrderUpdateRequest) GetRefundFee() int64 {
    return r.refundFee
}

func (r *TaobaoXhotelOrderUpdateRequest) SetReasonType(reasonType int64) error {
    r.reasonType = reasonType
    r.Set("reason_type", reasonType)
    return nil
}

func (r TaobaoXhotelOrderUpdateRequest) GetReasonType() int64 {
    return r.reasonType
}

func (r *TaobaoXhotelOrderUpdateRequest) SetInvoiceAmount(invoiceAmount int64) error {
    r.invoiceAmount = invoiceAmount
    r.Set("invoice_amount", invoiceAmount)
    return nil
}

func (r TaobaoXhotelOrderUpdateRequest) GetInvoiceAmount() int64 {
    return r.invoiceAmount
}

