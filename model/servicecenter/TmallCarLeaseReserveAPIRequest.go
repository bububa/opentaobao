package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseReserveAPIRequest 整车租车回传预约信息 API请求
// tmall.car.lease.reserve
//
// 租赁公司回传预约到店信息
type TmallCarLeaseReserveAPIRequest struct {
	model.Params
	// 文案
	_text string
	// 车架号
	_vin string
	// 买家昵称
	_buyerNick string
	// 买家id
	_buyerId int64
	// 订单id
	_orderId int64
	// 1 代表 车辆到店，已预约用户到店提车   ; 2 车辆到店，未能联系到用户
	_flag int64
}

// NewTmallCarLeaseReserveRequest 初始化TmallCarLeaseReserveAPIRequest对象
func NewTmallCarLeaseReserveRequest() *TmallCarLeaseReserveAPIRequest {
	return &TmallCarLeaseReserveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseReserveAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.reserve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseReserveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetText is Text Setter
// 文案
func (r *TmallCarLeaseReserveAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TmallCarLeaseReserveAPIRequest) GetText() string {
	return r._text
}

// SetVin is Vin Setter
// 车架号
func (r *TmallCarLeaseReserveAPIRequest) SetVin(_vin string) error {
	r._vin = _vin
	r.Set("vin", _vin)
	return nil
}

// GetVin Vin Getter
func (r TmallCarLeaseReserveAPIRequest) GetVin() string {
	return r._vin
}

// SetBuyerNick is BuyerNick Setter
// 买家昵称
func (r *TmallCarLeaseReserveAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r TmallCarLeaseReserveAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetBuyerId is BuyerId Setter
// 买家id
func (r *TmallCarLeaseReserveAPIRequest) SetBuyerId(_buyerId int64) error {
	r._buyerId = _buyerId
	r.Set("buyer_id", _buyerId)
	return nil
}

// GetBuyerId BuyerId Getter
func (r TmallCarLeaseReserveAPIRequest) GetBuyerId() int64 {
	return r._buyerId
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TmallCarLeaseReserveAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarLeaseReserveAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetFlag is Flag Setter
// 1 代表 车辆到店，已预约用户到店提车   ; 2 车辆到店，未能联系到用户
func (r *TmallCarLeaseReserveAPIRequest) SetFlag(_flag int64) error {
	r._flag = _flag
	r.Set("flag", _flag)
	return nil
}

// GetFlag Flag Getter
func (r TmallCarLeaseReserveAPIRequest) GetFlag() int64 {
	return r._flag
}
