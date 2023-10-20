package hotelhstdf

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripHotelHstdfShotelExportshotelAPIResponse 商家自主导出相似度高的标准酒店 API返回值
// alitrip.hotel.hstdf.shotel.exportshotel
//
// 商家通过给出自己的卖家酒店信息，通过接口可以返回相似度高的标准酒店信息
type AlitripHotelHstdfShotelExportshotelAPIResponse struct {
	model.CommonResponse
	AlitripHotelHstdfShotelExportshotelAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelExportshotelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripHotelHstdfShotelExportshotelAPIResponseModel).Reset()
}

// AlitripHotelHstdfShotelExportshotelAPIResponseModel is 商家自主导出相似度高的标准酒店 成功返回结果
type AlitripHotelHstdfShotelExportshotelAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_hotel_hstdf_shotel_exportshotel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// top返回结果
	TopResultSet *TopResultSet `json:"top_result_set,omitempty" xml:"top_result_set,omitempty"`
}

// Reset 清空结构体
func (m *AlitripHotelHstdfShotelExportshotelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopResultSet = nil
}

var poolAlitripHotelHstdfShotelExportshotelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripHotelHstdfShotelExportshotelAPIResponse)
	},
}

// GetAlitripHotelHstdfShotelExportshotelAPIResponse 从 sync.Pool 获取 AlitripHotelHstdfShotelExportshotelAPIResponse
func GetAlitripHotelHstdfShotelExportshotelAPIResponse() *AlitripHotelHstdfShotelExportshotelAPIResponse {
	return poolAlitripHotelHstdfShotelExportshotelAPIResponse.Get().(*AlitripHotelHstdfShotelExportshotelAPIResponse)
}

// ReleaseAlitripHotelHstdfShotelExportshotelAPIResponse 将 AlitripHotelHstdfShotelExportshotelAPIResponse 保存到 sync.Pool
func ReleaseAlitripHotelHstdfShotelExportshotelAPIResponse(v *AlitripHotelHstdfShotelExportshotelAPIResponse) {
	v.Reset()
	poolAlitripHotelHstdfShotelExportshotelAPIResponse.Put(v)
}
