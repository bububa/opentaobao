package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪门票核销通知 API请求
taobao.travel.ticket.order.verify

系统商通过TOP接口调用通知飞猪门票核销情况
*/
type TaobaoTravelTicketOrderVerifyRequest struct {
    model.Params
    // 下单订单ID
    orderId   int64
    // 外部订单ID
    outOrderId   string
    // 使用凭证信息
    voucherInfos   []VoucherInfoDto
    // 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
    writeOffType   int64
}

// 初始化TaobaoTravelTicketOrderVerifyRequest对象
func NewTaobaoTravelTicketOrderVerifyRequest() *TaobaoTravelTicketOrderVerifyRequest{
    return &TaobaoTravelTicketOrderVerifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTravelTicketOrderVerifyRequest) GetApiMethodName() string {
    return "taobao.travel.ticket.order.verify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTravelTicketOrderVerifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 下单订单ID
func (r *TaobaoTravelTicketOrderVerifyRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetOrderId() int64 {
    return r.orderId
}
// OutOrderId Setter
// 外部订单ID
func (r *TaobaoTravelTicketOrderVerifyRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetOutOrderId() string {
    return r.outOrderId
}
// VoucherInfos Setter
// 使用凭证信息
func (r *TaobaoTravelTicketOrderVerifyRequest) SetVoucherInfos(voucherInfos []VoucherInfoDto) error {
    r.voucherInfos = voucherInfos
    r.Set("voucher_infos", voucherInfos)
    return nil
}

// VoucherInfos Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetVoucherInfos() []VoucherInfoDto {
    return r.voucherInfos
}
// WriteOffType Setter
// 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
func (r *TaobaoTravelTicketOrderVerifyRequest) SetWriteOffType(writeOffType int64) error {
    r.writeOffType = writeOffType
    r.Set("write_off_type", writeOffType)
    return nil
}

// WriteOffType Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetWriteOffType() int64 {
    return r.writeOffType
}
