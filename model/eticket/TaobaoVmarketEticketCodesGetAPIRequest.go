package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketcodesgetAPIRequest 电子凭证码列表查询 API请求
// taobao.vmarket.eticket.codes.get
//
// 查询某个订单的所有码的列表
type TaobaovmarketeticketcodesgetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
	// 码商ID
	_codemerchantId int64
}

// NewTaobaovmarketeticketcodesgetRequest 初始化TaobaovmarketeticketcodesgetAPIRequest对象
func NewTaobaovmarketeticketcodesgetRequest() *TaobaovmarketeticketcodesgetAPIRequest {
	return &TaobaovmarketeticketcodesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketcodesgetAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.codes.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketcodesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketcodesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TaobaovmarketeticketcodesgetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovmarketeticketcodesgetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID
func (r *TaobaovmarketeticketcodesgetAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaovmarketeticketcodesgetAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
