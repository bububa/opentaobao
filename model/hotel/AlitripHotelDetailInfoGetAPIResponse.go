package hotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
详情页动态信息接口 API返回值 
alitrip.hotel.detail.info.get

酒店详情页动态信息TOP方法
*/
type AlitripHotelDetailInfoGetAPIResponse struct {
    model.CommonResponse
    AlitripHotelDetailInfoGetAPIResponseModel
}

// 详情页动态信息接口 成功返回结果
type AlitripHotelDetailInfoGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_hotel_detail_info_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlitripHotelDetailInfoGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
