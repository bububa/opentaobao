package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallalihousetradecouponordercodeexchangeAPIRequest 核销券码 API请求
// tmall.alihouse.trade.coupon.order.code.exchange
//
// ETC核销券码
type TmallalihousetradecouponordercodeexchangeAPIRequest struct {
	model.Params
	// 核销码
	_exchangeCodeDto *ExchangeCodeDto
}

// NewTmallalihousetradecouponordercodeexchangeRequest 初始化TmallalihousetradecouponordercodeexchangeAPIRequest对象
func NewTmallalihousetradecouponordercodeexchangeRequest() *TmallalihousetradecouponordercodeexchangeAPIRequest {
	return &TmallalihousetradecouponordercodeexchangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallalihousetradecouponordercodeexchangeAPIRequest) GetApiMethodName() string {
	return "tmall.alihouse.trade.coupon.order.code.exchange"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallalihousetradecouponordercodeexchangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallalihousetradecouponordercodeexchangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExchangeCodeDto is ExchangeCodeDto Setter
// 核销码
func (r *TmallalihousetradecouponordercodeexchangeAPIRequest) SetExchangeCodeDto(_exchangeCodeDto *ExchangeCodeDto) error {
	r._exchangeCodeDto = _exchangeCodeDto
	r.Set("exchange_code_dto", _exchangeCodeDto)
	return nil
}

// GetExchangeCodeDto ExchangeCodeDto Getter
func (r TmallalihousetradecouponordercodeexchangeAPIRequest) GetExchangeCodeDto() *ExchangeCodeDto {
	return r._exchangeCodeDto
}
