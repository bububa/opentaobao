package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票更新价格库存 APIResponse
taobao.bus.numbers.stockprice.update

用于汽车票代理商更新价格库存
*/
type TaobaoBusNumbersStockpriceUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusNumbersStockpriceUpdateResponse `json:"taobao_bus_numbers_stockprice_update_response,omitempty"`
}

type TaobaoBusNumbersStockpriceUpdateResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 成功数量
    SuccCount   int64 `json:"succ_count,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
