package hotelalliance

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelSingleInfoGetAPIResponse
获取单体酒店信息 API返回值
alitrip.hotel.single.info.get

用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务 */
type AlitripHotelSingleInfoGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelSingleInfoGetAPIResponseModel
}

// AlitripHotelSingleInfoGetAPIResponseModel is 获取单体酒店信息 成功返回结果
type AlitripHotelSingleInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_single_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回Result
	HmsTopResultSet *HmsTopResultSet `json:"hms_top_result_set,omitempty" xml:"hms_top_result_set,omitempty"`
}
