package hotelalliance

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取联盟hid APIResponse
alitrip.hotel.alliance.hid.get

获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询
*/
type AlitripHotelAllianceHidGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelAllianceHidGetResponse
}

type AlitripHotelAllianceHidGetResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_alliance_hid_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    HmsTopResultSet   *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`

    
}
