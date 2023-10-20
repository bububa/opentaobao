package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHtcouponFuluPhonechargeCallbackAPIRequest 话费充值回调 API请求
// alibaba.htcoupon.fulu.phonecharge.callback
//
// 话费充值为异步操作，此接口用于充值成功后，供应商回调。
type AlibabaHtcouponFuluPhonechargeCallbackAPIRequest struct {
	model.Params
	// 充值账号不是手机号
	_errorMessage string
	// 19062837751058701652
	_outOrderId string
	// 订单状态: (success:成功， processing:处理中，failed:失败， untreated:未处理)
	_orderState string
	// 欢行订单号
	_htOrderId string
}

// NewAlibabaHtcouponFuluPhonechargeCallbackRequest 初始化AlibabaHtcouponFuluPhonechargeCallbackAPIRequest对象
func NewAlibabaHtcouponFuluPhonechargeCallbackRequest() *AlibabaHtcouponFuluPhonechargeCallbackAPIRequest {
	return &AlibabaHtcouponFuluPhonechargeCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.htcoupon.fulu.phonecharge.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErrorMessage is ErrorMessage Setter
// 充值账号不是手机号
func (r *AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) SetErrorMessage(_errorMessage string) error {
	r._errorMessage = _errorMessage
	r.Set("error_message", _errorMessage)
	return nil
}

// GetErrorMessage ErrorMessage Getter
func (r AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) GetErrorMessage() string {
	return r._errorMessage
}

// SetOutOrderId is OutOrderId Setter
// 19062837751058701652
func (r *AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetOrderState is OrderState Setter
// 订单状态: (success:成功， processing:处理中，failed:失败， untreated:未处理)
func (r *AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) SetOrderState(_orderState string) error {
	r._orderState = _orderState
	r.Set("order_state", _orderState)
	return nil
}

// GetOrderState OrderState Getter
func (r AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) GetOrderState() string {
	return r._orderState
}

// SetHtOrderId is HtOrderId Setter
// 欢行订单号
func (r *AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) SetHtOrderId(_htOrderId string) error {
	r._htOrderId = _htOrderId
	r.Set("ht_order_id", _htOrderId)
	return nil
}

// GetHtOrderId HtOrderId Getter
func (r AlibabaHtcouponFuluPhonechargeCallbackAPIRequest) GetHtOrderId() string {
	return r._htOrderId
}
