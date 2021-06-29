package hotelhstdf

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据城市查询商圈 APIResponse
alitrip.hotel.hstdf.businessarea.get

根据cityId分页查询商圈信息
*/
type AlitripHotelHstdfBusinessareaGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelHstdfBusinessareaGetResponse
}

type AlitripHotelHstdfBusinessareaGetResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hstdf_businessarea_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // top返回结果
    
    Result   *TopStdResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
