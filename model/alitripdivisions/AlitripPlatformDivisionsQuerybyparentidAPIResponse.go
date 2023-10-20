package alitripdivisions

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformDivisionsQuerybyparentidAPIResponse 根据父节点id查询下级行政区划数据 API返回值
// alitrip.platform.divisions.querybyparentid
//
// 根据行政区划id查询下一层级行政区划数据
type AlitripPlatformDivisionsQuerybyparentidAPIResponse struct {
	model.CommonResponse
	AlitripPlatformDivisionsQuerybyparentidAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPlatformDivisionsQuerybyparentidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPlatformDivisionsQuerybyparentidAPIResponseModel).Reset()
}

// AlitripPlatformDivisionsQuerybyparentidAPIResponseModel is 根据父节点id查询下级行政区划数据 成功返回结果
type AlitripPlatformDivisionsQuerybyparentidAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_divisions_querybyparentid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripPlatformDivisionsQuerybyparentidResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPlatformDivisionsQuerybyparentidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPlatformDivisionsQuerybyparentidAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPlatformDivisionsQuerybyparentidAPIResponse)
	},
}

// GetAlitripPlatformDivisionsQuerybyparentidAPIResponse 从 sync.Pool 获取 AlitripPlatformDivisionsQuerybyparentidAPIResponse
func GetAlitripPlatformDivisionsQuerybyparentidAPIResponse() *AlitripPlatformDivisionsQuerybyparentidAPIResponse {
	return poolAlitripPlatformDivisionsQuerybyparentidAPIResponse.Get().(*AlitripPlatformDivisionsQuerybyparentidAPIResponse)
}

// ReleaseAlitripPlatformDivisionsQuerybyparentidAPIResponse 将 AlitripPlatformDivisionsQuerybyparentidAPIResponse 保存到 sync.Pool
func ReleaseAlitripPlatformDivisionsQuerybyparentidAPIResponse(v *AlitripPlatformDivisionsQuerybyparentidAPIResponse) {
	v.Reset()
	poolAlitripPlatformDivisionsQuerybyparentidAPIResponse.Put(v)
}
