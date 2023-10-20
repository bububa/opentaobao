package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest 核销券码 API请求
// tmall.alihouse.trade.coupon.order.code.exchange
//
// ETC核销券码
type TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest struct {
	model.Params
	// 核销码
	_exchangeCodeDto *ExchangeCodeDto
}

// NewTmallAlihouseTradeCouponOrderCodeExchangeRequest 初始化TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest对象
func NewTmallAlihouseTradeCouponOrderCodeExchangeRequest() *TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest {
	return &TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.code.exchange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExchangeCodeDto is ExchangeCodeDto Setter
// 核销码
func (r *TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest) SetExchangeCodeDto(_exchangeCodeDto *ExchangeCodeDto) error {
	r._exchangeCodeDto = _exchangeCodeDto
	r.Set("exchange_code_dto", _exchangeCodeDto)
	return nil
}

// GetExchangeCodeDto ExchangeCodeDto Getter
func (r TmallAlihouseTradeCouponOrderCodeExchangeAPIRequest) GetExchangeCodeDto() *ExchangeCodeDto {
	return r._exchangeCodeDto
}
