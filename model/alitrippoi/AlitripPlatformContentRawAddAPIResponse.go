package alitrippoi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPlatformContentRawAddAPIResponse 穷游内容写入接口 API返回值
// alitrip.platform.content.raw.add
//
// 穷游内容写入飞猪接口
type AlitripPlatformContentRawAddAPIResponse struct {
	model.CommonResponse
	AlitripPlatformContentRawAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripPlatformContentRawAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripPlatformContentRawAddAPIResponseModel).Reset()
}

// AlitripPlatformContentRawAddAPIResponseModel is 穷游内容写入接口 成功返回结果
type AlitripPlatformContentRawAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_platform_content_raw_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripPlatformContentRawAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripPlatformContentRawAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripPlatformContentRawAddAPIResponse)
	},
}

// GetAlitripPlatformContentRawAddAPIResponse 从 sync.Pool 获取 AlitripPlatformContentRawAddAPIResponse
func GetAlitripPlatformContentRawAddAPIResponse() *AlitripPlatformContentRawAddAPIResponse {
	return poolAlitripPlatformContentRawAddAPIResponse.Get().(*AlitripPlatformContentRawAddAPIResponse)
}

// ReleaseAlitripPlatformContentRawAddAPIResponse 将 AlitripPlatformContentRawAddAPIResponse 保存到 sync.Pool
func ReleaseAlitripPlatformContentRawAddAPIResponse(v *AlitripPlatformContentRawAddAPIResponse) {
	v.Reset()
	poolAlitripPlatformContentRawAddAPIResponse.Put(v)
}
