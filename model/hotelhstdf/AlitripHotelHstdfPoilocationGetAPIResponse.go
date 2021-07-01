package hotelhstdf

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfPoilocationGetAPIResponse
根据平台城市id分页查询poi location API返回值
alitrip.hotel.hstdf.poilocation.get

根据平台城市id分页查询poi location */
type AlitripHotelHstdfPoilocationGetAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfPoilocationGetAPIResponseModel
}

// AlitripHotelHstdfPoilocationGetAPIResponseModel is 根据平台城市id分页查询poi location 成功返回结果
type AlitripHotelHstdfPoilocationGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_poilocation_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	Result *TopStdResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
