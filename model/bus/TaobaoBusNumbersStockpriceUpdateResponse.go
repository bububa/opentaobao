package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票更新价格库存 APIResponse
taobao.bus.numbers.stockprice.update

用于汽车票代理商更新价格库存
*/
type TaobaoBusNumbersStockpriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_numbers_stockprice_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"