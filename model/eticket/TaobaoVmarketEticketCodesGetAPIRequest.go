package eticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketCodesGetAPIRequest 电子凭证码列表查询 API请求
// taobao.vmarket.eticket.codes.get
//
// 查询某个订单的所有码的列表
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketCodesGetAPIRequest) Reset() {
	r._orderId = 0
	r._codemerchantId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.codes.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TaobaoVmarketEticketCodesGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCodemerchantId is CodemerchantId Setter
// 码商ID
func (r *TaobaoVmarketEticketCodesGetAPIRequest) SetCodemerchantId(_codemerchantId int64) error {
	r._codemerchantId = _codemerchantId
	r.Set("codemerchant_id", _codemerchantId)
	return nil
}

// GetCodemerchantId CodemerchantId Getter
func (r TaobaoVmarketEticketCodesGetAPIRequest) GetCodemerchantId() int64 {
	return r._codemerchantId
}

var poolTaobaoVmarketEticketCodesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketCodesGetRequest()
	},
}

// GetTaobaoVmarketEticketCodesGetRequest 从 sync.Pool 获取 TaobaoVmarketEticketCodesGetAPIRequest
func GetTaobaoVmarketEticketCodesGetAPIRequest() *TaobaoVmarketEticketCodesGetAPIRequest {
	return poolTaobaoVmarketEticketCodesGetAPIRequest.Get().(*TaobaoVmarketEticketCodesGetAPIRequest)
}

// ReleaseTaobaoVmarketEticketCodesGetAPIRequest 将 TaobaoVmarketEticketCodesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketCodesGetAPIRequest(v *TaobaoVmarketEticketCodesGetAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketCodesGetAPIRequest.Put(v)
}
