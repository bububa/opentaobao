package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用车订单查询接口 APIResponse
alitrip.btrip.vehicle.order.search

企业获取商旅用车订单数据
*/
type AlitripBtripVehicleOrderSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripVehicleOrderSearchResponse
}

type AlitripBtripVehicleOrderSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_vehicle_order_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *BtriphomeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
