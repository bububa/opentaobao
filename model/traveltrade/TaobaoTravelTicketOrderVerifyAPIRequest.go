package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTravelTicketOrderVerifyAPIRequest 飞猪门票核销通知 API请求
// taobao.travel.ticket.order.verify
//
// 系统商通过TOP接口调用通知飞猪门票核销情况
type TaobaoTravelTicketOrderVerifyAPIRequest struct {
	model.Params
	// 下单订单ID
	_orderId int64
	// 外部订单ID
	_outOrderId string
	// 使用凭证信息
	_voucherInfos []VoucherInfoDto
	// 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
	_writeOffType int64
}

// NewTaobaoTravelTicketOrderVerifyRequest 初始化TaobaoTravelTicketOrderVerifyAPIRequest对象
func NewTaobaoTravelTicketOrderVerifyRequest() *TaobaoTravelTicketOrderVerifyAPIRequest {
	return &TaobaoTravelTicketOrderVerifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetApiMethodName() string {
	return "taobao.travel.ticket.order.verify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderId is OrderId Setter
// 下单订单ID
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetOutOrderId is OutOrderId Setter
// 外部订单ID
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetVoucherInfos is VoucherInfos Setter
// 使用凭证信息
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetVoucherInfos(_voucherInfos []VoucherInfoDto) error {
	r._voucherInfos = _voucherInfos
	r.Set("voucher_infos", _voucherInfos)
	return nil
}

// GetVoucherInfos VoucherInfos Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetVoucherInfos() []VoucherInfoDto {
	return r._voucherInfos
}

// SetWriteOffType is WriteOffType Setter
// 供应商核销回调类型：0表示使用本次核销数量（常规），1表示使用总核销数量（已使用+本次）
func (r *TaobaoTravelTicketOrderVerifyAPIRequest) SetWriteOffType(_writeOffType int64) error {
	r._writeOffType = _writeOffType
	r.Set("write_off_type", _writeOffType)
	return nil
}

// GetWriteOffType WriteOffType Getter
func (r TaobaoTravelTicketOrderVerifyAPIRequest) GetWriteOffType() int64 {
	return r._writeOffType
}
