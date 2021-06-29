package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】用车订单搜索 APIResponse
alitrip.btrip.supplychain.vehicle.search

【商旅】用车订单搜索
*/
type AlitripBtripSupplychainVehicleSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainVehicleSearchResponse
}

type AlitripBtripSupplychainVehicleSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_vehicle_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
