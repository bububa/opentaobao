package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketCodesGetAPIRequest
电子凭证码列表查询 API请求
taobao.vmarket.eticket.codes.get

查询某个订单的所有码的列表 */
type TaobaoVmarketEticketCodesGetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
	// 码商ID
	_codemerchantId int64
}

// NewTaobaoVmarketEticketCodesGetRequest 初始化TaobaoVmarketEticketCodesGetAPIRequest对象
func NewTaobaoVmarketEticketCodesGetRequest() *TaobaoVmarketEticketCodesGetAPIRequest {
	return &TaobaoVmarketEticketCodesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.codes.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单号
func (r *TaobaoVmarketEticketCodesGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is CodemerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketCodesGetAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// Get CodemerchantId Getter
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}
