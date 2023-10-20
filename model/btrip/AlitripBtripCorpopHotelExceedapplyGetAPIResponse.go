package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopHotelExceedapplyGetAPIResponse 商旅酒店第三方超标审批单搜索接口 API返回值
// alitrip.btrip.corpop.hotel.exceedapply.get
//
// 商旅酒店第三方超标审批单搜索接口
type AlitripBtripCorpopHotelExceedapplyGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopHotelExceedapplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopHotelExceedapplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopHotelExceedapplyGetAPIResponseModel).Reset()
}

// AlitripBtripCorpopHotelExceedapplyGetAPIResponseModel is 商旅酒店第三方超标审批单搜索接口 成功返回结果
type AlitripBtripCorpopHotelExceedapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_hotel_exceedapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopHotelExceedapplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopHotelExceedapplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopHotelExceedapplyGetAPIResponse)
	},
}

// GetAlitripBtripCorpopHotelExceedapplyGetAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopHotelExceedapplyGetAPIResponse
func GetAlitripBtripCorpopHotelExceedapplyGetAPIResponse() *AlitripBtripCorpopHotelExceedapplyGetAPIResponse {
	return poolAlitripBtripCorpopHotelExceedapplyGetAPIResponse.Get().(*AlitripBtripCorpopHotelExceedapplyGetAPIResponse)
}

// ReleaseAlitripBtripCorpopHotelExceedapplyGetAPIResponse 将 AlitripBtripCorpopHotelExceedapplyGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopHotelExceedapplyGetAPIResponse(v *AlitripBtripCorpopHotelExceedapplyGetAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopHotelExceedapplyGetAPIResponse.Put(v)
}
