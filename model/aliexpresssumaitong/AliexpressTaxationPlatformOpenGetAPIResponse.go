package aliexpresssumaitong

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTaxationPlatformOpenGetAPIResponse 平台固定参数获取 API返回值
// aliexpress.taxation.platform.open.get
//
// Aliexpress开放平台固定参数获取
type AliexpressTaxationPlatformOpenGetAPIResponse struct {
	model.CommonResponse
	AliexpressTaxationPlatformOpenGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressTaxationPlatformOpenGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressTaxationPlatformOpenGetAPIResponseModel).Reset()
}

// AliexpressTaxationPlatformOpenGetAPIResponseModel is 平台固定参数获取 成功返回结果
type AliexpressTaxationPlatformOpenGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_taxation_platform_open_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressTaxationPlatformOpenGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressTaxationPlatformOpenGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressTaxationPlatformOpenGetAPIResponse)
	},
}

// GetAliexpressTaxationPlatformOpenGetAPIResponse 从 sync.Pool 获取 AliexpressTaxationPlatformOpenGetAPIResponse
func GetAliexpressTaxationPlatformOpenGetAPIResponse() *AliexpressTaxationPlatformOpenGetAPIResponse {
	return poolAliexpressTaxationPlatformOpenGetAPIResponse.Get().(*AliexpressTaxationPlatformOpenGetAPIResponse)
}

// ReleaseAliexpressTaxationPlatformOpenGetAPIResponse 将 AliexpressTaxationPlatformOpenGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressTaxationPlatformOpenGetAPIResponse(v *AliexpressTaxationPlatformOpenGetAPIResponse) {
	v.Reset()
	poolAliexpressTaxationPlatformOpenGetAPIResponse.Put(v)
}
