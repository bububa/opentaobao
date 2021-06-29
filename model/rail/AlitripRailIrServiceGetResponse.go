package rail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国际火车票仓位坐席查询 APIResponse
alitrip.rail.ir.service.get

国际火车票标准仓位坐席查询
*/
type AlitripRailIrServiceGetAPIResponse struct {
    model.CommonResponse
    AlitripRailIrServiceGetResponse
}

type AlitripRailIrServiceGetResponse struct {
    XMLName xml.Name `xml:"alitrip_rail_ir_service_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlitripRailIrServiceGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
