package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票余票接口 APIResponse
taobao.bus.seatprice.get

提供给商家，查询汽车票班次余票
*/
type TaobaoBusSeatpriceGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusSeatpriceGetResponse `json:"bus_seatprice_get_response,omitempty"` 
    TaobaoBusSeatpriceGetResponse
}

/* model for simplify = false
type TaobaoBusSeatpriceGetResponse struct {

    // 返回结果
    
    Result  *struct {
        TaobaoBusSeatpriceGetResultSet  *TaobaoBusSeatpriceGetResultSet `json:"taobao_bus_seatprice_get_result_set,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBusSeatpriceGetResponse struct {

    // 返回结果
    Result   *TaobaoBusSeatpriceGetResultSet `json:"result,omitempty"`

}
