package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票更新价格库存 APIRequest
taobao.bus.numbers.stockprice.update

用于汽车票代理商更新价格库存
*/
type TaobaoBusNumbersStockpriceUpdateRequest struct {
    model.Params

    // 请求参数
    paramTopBusPriceAndStockUpdateRQ   *TopBusPriceAndStockUpdateRq 

}

func NewTaobaoBusNumbersStockpriceUpdateRequest() *TaobaoBusNumbersStockpriceUpdateRequest{
    return &TaobaoBusNumbersStockpriceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusNumbersStockpriceUpdateRequest) GetApiMethodName() string {
    return "taobao.bus.numbers.stockprice.update"
}

func (r TaobaoBusNumbersStockpriceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusNumbersStockpriceUpdateRequest) SetParamTopBusPriceAndStockUpdateRQ(paramTopBusPriceAndStockUpdateRQ *TopBusPriceAndStockUpdateRq) error {
    r.paramTopBusPriceAndStockUpdateRQ = paramTopBusPriceAndStockUpdateRQ
    r.Set("param_top_bus_price_and_stock_update_r_q", paramTopBusPriceAndStockUpdateRQ)
    return nil
}

func (r TaobaoBusNumbersStockpriceUpdateRequest) GetParamTopBusPriceAndStockUpdateRQ() *TopBusPriceAndStockUpdateRq {
    return r.paramTopBusPriceAndStockUpdateRQ
}

