package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】机票订单查询 APIResponse
alitrip.btrip.supplychain.flight.search

【商旅】机票订单查询
*/
type AlitripBtripSupplychainFlightSearchAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainFlightSearchResponse
}

type AlitripBtripSupplychainFlightSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_flight_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参对象
    
    Result   *HisvResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
