package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradediscountbillgetAPIRequest 订单优惠账单查询 API请求
// alibaba.wdk.trade.discount.bill.get
//
// 商家查询订单优惠账单
type AlibabawdktradediscountbillgetAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *OrderDiscountBillQueryRequest
}

// NewAlibabawdktradediscountbillgetRequest 初始化AlibabawdktradediscountbillgetAPIRequest对象
func NewAlibabawdktradediscountbillgetRequest() *AlibabawdktradediscountbillgetAPIRequest {
	return &AlibabawdktradediscountbillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktradediscountbillgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.discount.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktradediscountbillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktradediscountbillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabawdktradediscountbillgetAPIRequest) SetParam0(_param0 *OrderDiscountBillQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdktradediscountbillgetAPIRequest) GetParam0() *OrderDiscountBillQueryRequest {
	return r._param0
}
