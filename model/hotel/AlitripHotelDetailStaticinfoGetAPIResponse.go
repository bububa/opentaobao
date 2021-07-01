package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页静态信息 API返回值 
alitrip.hotel.detail.staticinfo.get

酒店静态信息TOP接口
*/
type AlitripHotelDetailStaticinfoGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelDetailStaticinfoGetAPIResponseModel
}

// 详情页静态信息 成功返回结果
type AlitripHotelDetailStaticinfoGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_hotel_detail_staticinfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlitripHotelDetailStaticinfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
