package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseCategoryControlSyncAPIResponse 类目权限上翻 API返回值
// alibaba.alihouse.category.control.sync
//
// 类目权限上翻
type AlibabaAlihouseCategoryControlSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseCategoryControlSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseCategoryControlSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseCategoryControlSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseCategoryControlSyncAPIResponseModel is 类目权限上翻 成功返回结果
type AlibabaAlihouseCategoryControlSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_category_control_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// list
	Data []CategoryControlResultDto `json:"data,omitempty" xml:"data>category_control_result_dto,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// message
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// success
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseCategoryControlSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseCategoryControlSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseCategoryControlSyncAPIResponse)
	},
}

// GetAlibabaAlihouseCategoryControlSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseCategoryControlSyncAPIResponse
func GetAlibabaAlihouseCategoryControlSyncAPIResponse() *AlibabaAlihouseCategoryControlSyncAPIResponse {
	return poolAlibabaAlihouseCategoryControlSyncAPIResponse.Get().(*AlibabaAlihouseCategoryControlSyncAPIResponse)
}

// ReleaseAlibabaAlihouseCategoryControlSyncAPIResponse 将 AlibabaAlihouseCategoryControlSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseCategoryControlSyncAPIResponse(v *AlibabaAlihouseCategoryControlSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseCategoryControlSyncAPIResponse.Put(v)
}
