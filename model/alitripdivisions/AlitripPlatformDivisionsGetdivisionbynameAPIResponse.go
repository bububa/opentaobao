package alitripdivisions

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformDivisionsGetdivisionbynameAPIResponse 根据中文名称与行政区划级别查询行政区划数据 API返回值
// alitrip.platform.divisions.getdivisionbyname
//
// 根据中文名称与行政区划级别查询行政区划数据
type AlitripPlatformDivisionsGetdivisionbynameAPIResponse struct {
	model.CommonResponse
	AlitripPlatformDivisionsGetdivisionbynameAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPlatformDivisionsGetdivisionbynameAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPlatformDivisionsGetdivisionbynameAPIResponseModel).Reset()
}

// AlitripPlatformDivisionsGetdivisionbynameAPIResponseModel is 根据中文名称与行政区划级别查询行政区划数据 成功返回结果
type AlitripPlatformDivisionsGetdivisionbynameAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_divisions_getdivisionbyname_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPlatformDivisionsGetdivisionbynameAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolAlitripPlatformDivisionsGetdivisionbynameAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPlatformDivisionsGetdivisionbynameAPIResponse)
	},
}

// GetAlitripPlatformDivisionsGetdivisionbynameAPIResponse 从 sync.Pool 获取 AlitripPlatformDivisionsGetdivisionbynameAPIResponse
func GetAlitripPlatformDivisionsGetdivisionbynameAPIResponse() *AlitripPlatformDivisionsGetdivisionbynameAPIResponse {
	return poolAlitripPlatformDivisionsGetdivisionbynameAPIResponse.Get().(*AlitripPlatformDivisionsGetdivisionbynameAPIResponse)
}

// ReleaseAlitripPlatformDivisionsGetdivisionbynameAPIResponse 将 AlitripPlatformDivisionsGetdivisionbynameAPIResponse 保存到 sync.Pool
func ReleaseAlitripPlatformDivisionsGetdivisionbynameAPIResponse(v *AlitripPlatformDivisionsGetdivisionbynameAPIResponse) {
	v.Reset()
	poolAlitripPlatformDivisionsGetdivisionbynameAPIResponse.Put(v)
}
