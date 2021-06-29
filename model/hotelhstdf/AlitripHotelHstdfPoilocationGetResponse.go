package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据平台城市id分页查询poi location APIResponse
alitrip.hotel.hstdf.poilocation.get

根据平台城市id分页查询poi location
*/
type AlitripHotelHstdfPoilocationGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfPoilocationGetResponse
}

type AlitripHotelHstdfPoilocationGetResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_poilocation_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // top返回结果
    
    Result   *TopStdResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
