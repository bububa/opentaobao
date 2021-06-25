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
    Response *TaobaoBusLastplaceGetResponse `json:"taobao_bus_lastplace_get_response,omitempty"`
}

type TaobaoBusLastplaceGetResponse struct {

    // 目的地返回结果
    Result   *TaobaoBusLastplaceGetResult `json:"result,omitempty"`

}
