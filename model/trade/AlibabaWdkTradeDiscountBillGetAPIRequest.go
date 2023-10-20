package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeDiscountBillGetAPIRequest 订单优惠账单查询 API请求
// alibaba.wdk.trade.discount.bill.get
//
// 商家查询订单优惠账单
type AlibabaWdkTradeDiscountBillGetAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *OrderDiscountBillQueryRequest
}

// NewAlibabaWdkTradeDiscountBillGetRequest 初始化AlibabaWdkTradeDiscountBillGetAPIRequest对象
func NewAlibabaWdkTradeDiscountBillGetRequest() *AlibabaWdkTradeDiscountBillGetAPIRequest {
	return &AlibabaWdkTradeDiscountBillGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeDiscountBillGetAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeDiscountBillGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.discount.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeDiscountBillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeDiscountBillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaWdkTradeDiscountBillGetAPIRequest) SetParam0(_param0 *OrderDiscountBillQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkTradeDiscountBillGetAPIRequest) GetParam0() *OrderDiscountBillQueryRequest {
	return r._param0
}

var poolAlibabaWdkTradeDiscountBillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeDiscountBillGetRequest()
	},
}

// GetAlibabaWdkTradeDiscountBillGetRequest 从 sync.Pool 获取 AlibabaWdkTradeDiscountBillGetAPIRequest
func GetAlibabaWdkTradeDiscountBillGetAPIRequest() *AlibabaWdkTradeDiscountBillGetAPIRequest {
	return poolAlibabaWdkTradeDiscountBillGetAPIRequest.Get().(*AlibabaWdkTradeDiscountBillGetAPIRequest)
}

// ReleaseAlibabaWdkTradeDiscountBillGetAPIRequest 将 AlibabaWdkTradeDiscountBillGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeDiscountBillGetAPIRequest(v *AlibabaWdkTradeDiscountBillGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeDiscountBillGetAPIRequest.Put(v)
}
