package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeMetroSyncAPIResponse 地铁数据同步 API返回值
// alibaba.alihouse.newhome.metro.sync
//
// 地铁数据同步
type AlibabaAlihouseNewhomeMetroSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeMetroSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeMetroSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeMetroSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeMetroSyncAPIResponseModel is 地铁数据同步 成功返回结果
type AlibabaAlihouseNewhomeMetroSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_metro_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeMetroSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeMetroSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeMetroSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeMetroSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeMetroSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeMetroSyncAPIResponse
func GetAlibabaAlihouseNewhomeMetroSyncAPIResponse() *AlibabaAlihouseNewhomeMetroSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeMetroSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeMetroSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeMetroSyncAPIResponse 将 AlibabaAlihouseNewhomeMetroSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeMetroSyncAPIResponse(v *AlibabaAlihouseNewhomeMetroSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeMetroSyncAPIResponse.Put(v)
}
