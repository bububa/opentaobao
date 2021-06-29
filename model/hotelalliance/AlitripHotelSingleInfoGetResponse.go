package hotelalliance

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单体酒店信息 APIResponse
alitrip.hotel.single.info.get

用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务
*/
type AlitripHotelSingleInfoGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelSingleInfoGetResponse
}

type AlitripHotelSingleInfoGetResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_single_info_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回Result
    
    HmsTopResultSet   *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`

    
}
