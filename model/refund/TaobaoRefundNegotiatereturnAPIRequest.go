package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundNegotiatereturnAPIRequest 协商退货退款 API请求
// taobao.refund.negotiatereturn
//
// 协商退货退款
type TaobaoRefundNegotiatereturnAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 退款版本号
	_refundVersion int64
	// 退款金额，单位（分）
	_refundFee int64
	// 地址ID，通过协商退货退款渲染接口获取到的
	_addressId int64
}

// NewTaobaoRefundNegotiatereturnRequest 初始化TaobaoRefundNegotiatereturnAPIRequest对象
func NewTaobaoRefundNegotiatereturnRequest() *TaobaoRefundNegotiatereturnAPIRequest {
	return &TaobaoRefundNegotiatereturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRefundNegotiatereturnAPIRequest) GetApiMethodName() string {
	return "taobao.refund.negotiatereturn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRefundNegotiatereturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaoRefundNegotiatereturnAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRefundNegotiatereturnAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundVersion is RefundVersion Setter
// 退款版本号
func (r *TaobaoRefundNegotiatereturnAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaoRefundNegotiatereturnAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}

// SetRefundFee is RefundFee Setter
// 退款金额，单位（分）
func (r *TaobaoRefundNegotiatereturnAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaoRefundNegotiatereturnAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetAddressId is AddressId Setter
// 地址ID，通过协商退货退款渲染接口获取到的
func (r *TaobaoRefundNegotiatereturnAPIRequest) SetAddressId(_addressId int64) error {
	r._addressId = _addressId
	r.Set("address_id", _addressId)
	return nil
}

// GetAddressId AddressId Getter
func (r TaobaoRefundNegotiatereturnAPIRequest) GetAddressId() int64 {
	return r._addressId
}
