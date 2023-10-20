package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseCategoryControlQueryAPIResponse 类目权限查询 API返回值
// alibaba.alihouse.category.control.query
//
// 类目权限查询
type AlibabaAlihouseCategoryControlQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseCategoryControlQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseCategoryControlQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseCategoryControlQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseCategoryControlQueryAPIResponseModel is 类目权限查询 成功返回结果
type AlibabaAlihouseCategoryControlQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_category_control_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data []CategoryControlDto `json:"data,omitempty" xml:"data>category_control_dto,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// message
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// success
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseCategoryControlQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseCategoryControlQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseCategoryControlQueryAPIResponse)
	},
}

// GetAlibabaAlihouseCategoryControlQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseCategoryControlQueryAPIResponse
func GetAlibabaAlihouseCategoryControlQueryAPIResponse() *AlibabaAlihouseCategoryControlQueryAPIResponse {
	return poolAlibabaAlihouseCategoryControlQueryAPIResponse.Get().(*AlibabaAlihouseCategoryControlQueryAPIResponse)
}

// ReleaseAlibabaAlihouseCategoryControlQueryAPIResponse 将 AlibabaAlihouseCategoryControlQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseCategoryControlQueryAPIResponse(v *AlibabaAlihouseCategoryControlQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseCategoryControlQueryAPIResponse.Put(v)
}
