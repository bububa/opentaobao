package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
机场数据查询 APIResponse
alitrip.btrip.supplychain.flight.city

机场数据查询
*/
type AlitripBtripSupplychainFlightCityAPIResponse struct {
    model.CommonResponse
    AlitripBtripSupplychainFlightCityResponse
}

type AlitripBtripSupplychainFlightCityResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_supplychain_flight_city_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果对象
    
    Result   *SuggestRs `json:"result,omitempty" xml:"result,omitempty"`

    
    // 结果码
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 结果信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
