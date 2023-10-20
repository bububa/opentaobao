package aliexpresssumaitong

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTaxationCalculateOpenQueryAPIResponse 关务所需的申报清关字段 API返回值
// aliexpress.taxation.calculate.open.query
//
// 关务所需的申报清关字段
type AliexpressTaxationCalculateOpenQueryAPIResponse struct {
	model.CommonResponse
	AliexpressTaxationCalculateOpenQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressTaxationCalculateOpenQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressTaxationCalculateOpenQueryAPIResponseModel).Reset()
}

// AliexpressTaxationCalculateOpenQueryAPIResponseModel is 关务所需的申报清关字段 成功返回结果
type AliexpressTaxationCalculateOpenQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_taxation_calculate_open_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressTaxationCalculateOpenQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressTaxationCalculateOpenQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressTaxationCalculateOpenQueryAPIResponse)
	},
}

// GetAliexpressTaxationCalculateOpenQueryAPIResponse 从 sync.Pool 获取 AliexpressTaxationCalculateOpenQueryAPIResponse
func GetAliexpressTaxationCalculateOpenQueryAPIResponse() *AliexpressTaxationCalculateOpenQueryAPIResponse {
	return poolAliexpressTaxationCalculateOpenQueryAPIResponse.Get().(*AliexpressTaxationCalculateOpenQueryAPIResponse)
}

// ReleaseAliexpressTaxationCalculateOpenQueryAPIResponse 将 AliexpressTaxationCalculateOpenQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressTaxationCalculateOpenQueryAPIResponse(v *AliexpressTaxationCalculateOpenQueryAPIResponse) {
	v.Reset()
	poolAliexpressTaxationCalculateOpenQueryAPIResponse.Put(v)
}
