package hotelalliance

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取合作商信息 API返回值 
alitrip.hotel.hms.partner.info.get

用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务
*/
type AlitripHotelHmsPartnerInfoGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelHmsPartnerInfoGetResponse
}

// 获取合作商信息 成功返回结果
type AlitripHotelHmsPartnerInfoGetResponse struct {
    XMLName xml.Name `xml:"alitrip_hotel_hms_partner_info_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回Result
    HmsTopResultSet   *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`
}
