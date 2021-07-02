package hotelhstdf

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelMatchshotelselfAPIResponse 自主匹配标准酒店以及卖家酒店 API返回值
// alitrip.hotel.hstdf.shotel.matchshotelself
//
// 商家通过指定的标准酒店id和卖家酒店id进行匹配
type AlitripHotelHstdfShotelMatchshotelselfAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfShotelMatchshotelselfAPIResponseModel
}

// AlitripHotelHstdfShotelMatchshotelselfAPIResponseModel is 自主匹配标准酒店以及卖家酒店 成功返回结果
type AlitripHotelHstdfShotelMatchshotelselfAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_matchshotelself_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// 错误信息
	Errormsg string `json:"errormsg,omitempty" xml:"errormsg,omitempty"`
	// 是否成功
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
