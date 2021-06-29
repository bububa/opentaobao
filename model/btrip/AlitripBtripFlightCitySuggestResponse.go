package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
机票城市搜索 APIResponse
alitrip.btrip.flight.city.suggest

提供机票城市搜索接口，提高OA用户对接效率
*/
type AlitripBtripFlightCitySuggestAPIResponse struct {
    model.CommonResponse
    AlitripBtripFlightCitySuggestResponse
}

type AlitripBtripFlightCitySuggestResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_flight_city_suggest_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *BtripApplyResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
