package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取目的地数据 APIResponse
taobao.bus.lastplace.get

传入城市 获取对应的目的地
*/
type TaobaoBusLastplaceGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusLastplaceGetResponse `json:"bus_lastplace_get_response,omitempty"` 
    TaobaoBusLastplaceGetResponse
}

/* model for simplify = false
type TaobaoBusLastplaceGetResponse struct {

    // 目的地返回结果
    
    Result  *struct {
        TaobaoBusLastplaceGetResult  *TaobaoBusLastplaceGetResult `json:"taobao_bus_lastplace_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBusLastplaceGetResponse struct {

    // 目的地返回结果
    Result   *TaobaoBusLastplaceGetResult `json:"result,omitempty"`

}
