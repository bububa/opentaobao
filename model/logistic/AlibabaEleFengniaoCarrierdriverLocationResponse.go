package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询骑手当前位置 APIResponse
alibaba.ele.fengniao.carrierdriver.location

查询骑手当前位置
*/
type AlibabaEleFengniaoCarrierdriverLocationAPIResponse struct {
    model.CommonResponse
    AlibabaEleFengniaoCarrierdriverLocationResponse
}

type AlibabaEleFengniaoCarrierdriverLocationResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_fengniao_carrierdriver_location_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // location
    
    Location   *Location `json:"location,omitempty" xml:"location,omitempty"`

    
    // 运单状态
    
    State   int64 `json:"state,omitempty" xml:"state,omitempty"`

    
    // 骑手电话
    
    CarrierDriverPhone   string `json:"carrier_driver_phone,omitempty" xml:"carrier_driver_phone,omitempty"`

    
    // 骑手姓名
    
    CarrierDriverName   string `json:"carrier_driver_name,omitempty" xml:"carrier_driver_name,omitempty"`

    
    // 运单状态变化时间点
    
    OccurredAt   int64 `json:"occurred_at,omitempty" xml:"occurred_at,omitempty"`

    
}
