package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusNumbersStockpriceUpdateAPIRequest 汽车票更新价格库存 API请求
// taobao.bus.numbers.stockprice.update
//
// 用于汽车票代理商更新价格库存
type TaobaoBusNumbersStockpriceUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramTopBusPriceAndStockUpdateRQ *TopBusPriceAndStockUpdateRq
}

// NewTaobaoBusNumbersStockpriceUpdateRequest 初始化TaobaoBusNumbersStockpriceUpdateAPIRequest对象
func NewTaobaoBusNumbersStockpriceUpdateRequest() *TaobaoBusNumbersStockpriceUpdateAPIRequest {
	return &TaobaoBusNumbersStockpriceUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusNumbersStockpriceUpdateAPIRequest) Reset() {
	r._paramTopBusPriceAndStockUpdateRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusNumbersStockpriceUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.bus.numbers.stockprice.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusNumbersStockpriceUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusNumbersStockpriceUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamTopBusPriceAndStockUpdateRQ is ParamTopBusPriceAndStockUpdateRQ Setter
// 请求参数
func (r *TaobaoBusNumbersStockpriceUpdateAPIRequest) SetParamTopBusPriceAndStockUpdateRQ(_paramTopBusPriceAndStockUpdateRQ *TopBusPriceAndStockUpdateRq) error {
	r._paramTopBusPriceAndStockUpdateRQ = _paramTopBusPriceAndStockUpdateRQ
	r.Set("param_top_bus_price_and_stock_update_r_q", _paramTopBusPriceAndStockUpdateRQ)
	return nil
}

// GetParamTopBusPriceAndStockUpdateRQ ParamTopBusPriceAndStockUpdateRQ Getter
func (r TaobaoBusNumbersStockpriceUpdateAPIRequest) GetParamTopBusPriceAndStockUpdateRQ() *TopBusPriceAndStockUpdateRq {
	return r._paramTopBusPriceAndStockUpdateRQ
}

var poolTaobaoBusNumbersStockpriceUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusNumbersStockpriceUpdateRequest()
	},
}

// GetTaobaoBusNumbersStockpriceUpdateRequest 从 sync.Pool 获取 TaobaoBusNumbersStockpriceUpdateAPIRequest
func GetTaobaoBusNumbersStockpriceUpdateAPIRequest() *TaobaoBusNumbersStockpriceUpdateAPIRequest {
	return poolTaobaoBusNumbersStockpriceUpdateAPIRequest.Get().(*TaobaoBusNumbersStockpriceUpdateAPIRequest)
}

// ReleaseTaobaoBusNumbersStockpriceUpdateAPIRequest 将 TaobaoBusNumbersStockpriceUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusNumbersStockpriceUpdateAPIRequest(v *TaobaoBusNumbersStockpriceUpdateAPIRequest) {
	v.Reset()
	poolTaobaoBusNumbersStockpriceUpdateAPIRequest.Put(v)
}
