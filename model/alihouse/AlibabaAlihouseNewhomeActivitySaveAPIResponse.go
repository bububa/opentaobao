package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeActivitySaveAPIResponse 新增或者更新销售活动 API返回值
// alibaba.alihouse.newhome.activity.save
//
// 新增或者更新销售活动
type AlibabaAlihouseNewhomeActivitySaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeActivitySaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeActivitySaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeActivitySaveAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeActivitySaveAPIResponseModel is 新增或者更新销售活动 成功返回结果
type AlibabaAlihouseNewhomeActivitySaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeActivitySaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeActivitySaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeActivitySaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivitySaveAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeActivitySaveAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeActivitySaveAPIResponse
func GetAlibabaAlihouseNewhomeActivitySaveAPIResponse() *AlibabaAlihouseNewhomeActivitySaveAPIResponse {
	return poolAlibabaAlihouseNewhomeActivitySaveAPIResponse.Get().(*AlibabaAlihouseNewhomeActivitySaveAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeActivitySaveAPIResponse 将 AlibabaAlihouseNewhomeActivitySaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeActivitySaveAPIResponse(v *AlibabaAlihouseNewhomeActivitySaveAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeActivitySaveAPIResponse.Put(v)
}
