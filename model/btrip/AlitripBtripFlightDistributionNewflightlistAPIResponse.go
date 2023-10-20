package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionNewflightlistAPIResponse 商旅机票航班列表接口，用于分销询价V2 API返回值
// alitrip.btrip.flight.distribution.newflightlist
//
// 商旅机票航班列表接口，用于分销询价V2
type AlitripBtripFlightDistributionNewflightlistAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionNewflightlistAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionNewflightlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripFlightDistributionNewflightlistAPIResponseModel).Reset()
}

// AlitripBtripFlightDistributionNewflightlistAPIResponseModel is 商旅机票航班列表接口，用于分销询价V2 成功返回结果
type AlitripBtripFlightDistributionNewflightlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_newflightlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应信息
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripFlightDistributionNewflightlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripFlightDistributionNewflightlistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripFlightDistributionNewflightlistAPIResponse)
	},
}

// GetAlitripBtripFlightDistributionNewflightlistAPIResponse 从 sync.Pool 获取 AlitripBtripFlightDistributionNewflightlistAPIResponse
func GetAlitripBtripFlightDistributionNewflightlistAPIResponse() *AlitripBtripFlightDistributionNewflightlistAPIResponse {
	return poolAlitripBtripFlightDistributionNewflightlistAPIResponse.Get().(*AlitripBtripFlightDistributionNewflightlistAPIResponse)
}

// ReleaseAlitripBtripFlightDistributionNewflightlistAPIResponse 将 AlitripBtripFlightDistributionNewflightlistAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripFlightDistributionNewflightlistAPIResponse(v *AlitripBtripFlightDistributionNewflightlistAPIResponse) {
	v.Reset()
	poolAlitripBtripFlightDistributionNewflightlistAPIResponse.Put(v)
}
