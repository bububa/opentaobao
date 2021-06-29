package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票更新价格库存 API请求
taobao.bus.numbers.stockprice.update

用于汽车票代理商更新价格库存
*/
type TaobaoBusNumbersStockpriceUpdateRequest struct {
    model.Params
    // 请求参数
    paramTopBusPriceAndStockUpdateRQ   *TopBusPriceAndStockUpdateRq
}

// 初始化TaobaoBusNumbersStockpriceUpdateRequest对象
func NewTaobaoBusNumbersStockpriceUpdateRequest() *TaobaoBusNumbersStockpriceUpdateRequest{
    return &TaobaoBusNumbersStockpriceUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusNumbersStockpriceUpdateRequest) GetApiMethodName() string {
    return "taobao.bus.numbers.stockprice.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusNumbersStockpriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopBusPriceAndStockUpdateRQ Setter
// 请求参数
func (r *TaobaoBusNumbersStockpriceUpdateRequest) SetParamTopBusPriceAndStockUpdateRQ(paramTopBusPriceAndStockUpdateRQ *TopBusPriceAndStockUpdateRq) error {
    r.paramTopBusPriceAndStockUpdateRQ = paramTopBusPriceAndStockUpdateRQ
    r.Set("param_top_bus_price_and_stock_update_r_q", paramTopBusPriceAndStockUpdateRQ)
    return nil
}

// ParamTopBusPriceAndStockUpdateRQ Getter
func (r TaobaoBusNumbersStockpriceUpdateRequest) GetParamTopBusPriceAndStockUpdateRQ() *TopBusPriceAndStockUpdateRq {
    return r.paramTopBusPriceAndStockUpdateRQ
}
