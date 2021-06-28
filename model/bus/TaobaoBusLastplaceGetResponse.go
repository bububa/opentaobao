package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取目的地数据 APIResponse
taobao.bus.lastplace.get

传入城市 获取对应的目的地
*/
type TaobaoBusLastplaceGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_lastplace_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 目的地返回结果
    
    Result   *TaobaoBusLastplaceGetResult `json:"result,omitempty" xml:"