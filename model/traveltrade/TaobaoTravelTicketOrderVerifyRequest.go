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
    _orderId   int64
    // 外部订单ID
    _outOrderId   string
    // 使用凭证信息
    _voucherInfos   []VoucherInfoDto
    // 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
    _writeOffType   int64
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
func (r *TaobaoTravelTicketOrderVerifyRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetOrderId() int64 {
    return r._orderId
}
// OutOrderId Setter
// 外部订单ID
func (r *TaobaoTravelTicketOrderVerifyRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetOutOrderId() string {
    return r._outOrderId
}
// VoucherInfos Setter
// 使用凭证信息
func (r *TaobaoTravelTicketOrderVerifyRequest) SetVoucherInfos(_voucherInfos []VoucherInfoDto) error {
    r._voucherInfos = _voucherInfos
    r.Set("voucher_infos", _voucherInfos)
    return nil
}

// VoucherInfos Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetVoucherInfos() []VoucherInfoDto {
    return r._voucherInfos
}
// WriteOffType Setter
// 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
func (r *TaobaoTravelTicketOrderVerifyRequest) SetWriteOffType(_writeOffType int64) error {
    r._writeOffType = _writeOffType
    r.Set("write_off_type", _writeOffType)
    return nil
}

// WriteOffType Getter
func (r TaobaoTravelTicketOrderVerifyRequest) GetWriteOffType() int64 {
    return r._writeOffType
}
