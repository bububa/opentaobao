package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseProjectActivitySyncAPIResponse 电商券数据同步 API返回值
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
type AlibabaAlihouseProjectActivitySyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseProjectActivitySyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseProjectActivitySyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseProjectActivitySyncAPIResponseModel).Reset()
}

// AlibabaAlihouseProjectActivitySyncAPIResponseModel is 电商券数据同步 成功返回结果
type AlibabaAlihouseProjectActivitySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_project_activity_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseProjectActivitySyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseProjectActivitySyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseProjectActivitySyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseProjectActivitySyncAPIResponse)
	},
}

// GetAlibabaAlihouseProjectActivitySyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseProjectActivitySyncAPIResponse
func GetAlibabaAlihouseProjectActivitySyncAPIResponse() *AlibabaAlihouseProjectActivitySyncAPIResponse {
	return poolAlibabaAlihouseProjectActivitySyncAPIResponse.Get().(*AlibabaAlihouseProjectActivitySyncAPIResponse)
}

// ReleaseAlibabaAlihouseProjectActivitySyncAPIResponse 将 AlibabaAlihouseProjectActivitySyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseProjectActivitySyncAPIResponse(v *AlibabaAlihouseProjectActivitySyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseProjectActivitySyncAPIResponse.Put(v)
}
