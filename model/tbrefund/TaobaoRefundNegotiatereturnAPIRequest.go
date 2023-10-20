package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundnegotiatereturnAPIRequest 协商退货退款 API请求
// taobao.refund.negotiatereturn
//
// 协商退货退款
type TaobaorefundnegotiatereturnAPIRequest struct {
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

// NewTaobaorefundnegotiatereturnRequest 初始化TaobaorefundnegotiatereturnAPIRequest对象
func NewTaobaorefundnegotiatereturnRequest() *TaobaorefundnegotiatereturnAPIRequest {
	return &TaobaorefundnegotiatereturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorefundnegotiatereturnAPIRequest) GetApiMethodName() string {
	return "taobao.refund.negotiatereturn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorefundnegotiatereturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorefundnegotiatereturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaorefundnegotiatereturnAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorefundnegotiatereturnAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundVersion is RefundVersion Setter
// 退款版本号
func (r *TaobaorefundnegotiatereturnAPIRequest) SetRefundVersion(_refundVersion int64) error {
	r._refundVersion = _refundVersion
	r.Set("refund_version", _refundVersion)
	return nil
}

// GetRefundVersion RefundVersion Getter
func (r TaobaorefundnegotiatereturnAPIRequest) GetRefundVersion() int64 {
	return r._refundVersion
}

// SetRefundFee is RefundFee Setter
// 退款金额，单位（分）
func (r *TaobaorefundnegotiatereturnAPIRequest) SetRefundFee(_refundFee int64) error {
	r._refundFee = _refundFee
	r.Set("refund_fee", _refundFee)
	return nil
}

// GetRefundFee RefundFee Getter
func (r TaobaorefundnegotiatereturnAPIRequest) GetRefundFee() int64 {
	return r._refundFee
}

// SetAddressId is AddressId Setter
// 地址ID，通过协商退货退款渲染接口获取到的
func (r *TaobaorefundnegotiatereturnAPIRequest) SetAddressId(_addressId int64) error {
	r._addressId = _addressId
	r.Set("address_id", _addressId)
	return nil
}

// GetAddressId AddressId Getter
func (r TaobaorefundnegotiatereturnAPIRequest) GetAddressId() int64 {
	return r._addressId
}
