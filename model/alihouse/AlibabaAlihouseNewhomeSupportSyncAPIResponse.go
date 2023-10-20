package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeSupportSyncAPIResponse 周边配套数据同步 API返回值
// alibaba.alihouse.newhome.support.sync
//
// 周边配套数据同步
type AlibabaAlihouseNewhomeSupportSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeSupportSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeSupportSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeSupportSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeSupportSyncAPIResponseModel is 周边配套数据同步 成功返回结果
type AlibabaAlihouseNewhomeSupportSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_support_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeSupportSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeSupportSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeSupportSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeSupportSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeSupportSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeSupportSyncAPIResponse
func GetAlibabaAlihouseNewhomeSupportSyncAPIResponse() *AlibabaAlihouseNewhomeSupportSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeSupportSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeSupportSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeSupportSyncAPIResponse 将 AlibabaAlihouseNewhomeSupportSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeSupportSyncAPIResponse(v *AlibabaAlihouseNewhomeSupportSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeSupportSyncAPIResponse.Put(v)
}
