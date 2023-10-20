package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopFlightExceedapplyGetAPIResponse 商旅机票第三方超标审批单搜索接口 API返回值
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
type AlitripBtripCorpopFlightExceedapplyGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripCorpopFlightExceedapplyGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripCorpopFlightExceedapplyGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripCorpopFlightExceedapplyGetAPIResponseModel).Reset()
}

// AlitripBtripCorpopFlightExceedapplyGetAPIResponseModel is 商旅机票第三方超标审批单搜索接口 成功返回结果
type AlitripBtripCorpopFlightExceedapplyGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_corpop_flight_exceedapply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripCorpopFlightExceedapplyGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripCorpopFlightExceedapplyGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripCorpopFlightExceedapplyGetAPIResponse)
	},
}

// GetAlitripBtripCorpopFlightExceedapplyGetAPIResponse 从 sync.Pool 获取 AlitripBtripCorpopFlightExceedapplyGetAPIResponse
func GetAlitripBtripCorpopFlightExceedapplyGetAPIResponse() *AlitripBtripCorpopFlightExceedapplyGetAPIResponse {
	return poolAlitripBtripCorpopFlightExceedapplyGetAPIResponse.Get().(*AlitripBtripCorpopFlightExceedapplyGetAPIResponse)
}

// ReleaseAlitripBtripCorpopFlightExceedapplyGetAPIResponse 将 AlitripBtripCorpopFlightExceedapplyGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripCorpopFlightExceedapplyGetAPIResponse(v *AlitripBtripCorpopFlightExceedapplyGetAPIResponse) {
	v.Reset()
	poolAlitripBtripCorpopFlightExceedapplyGetAPIResponse.Put(v)
}
