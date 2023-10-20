package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLineSyncAPIResponse 环线数据同步 API返回值
// alibaba.alihouse.newhome.line.sync
//
// 环线数据同步
type AlibabaAlihouseNewhomeLineSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeLineSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLineSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeLineSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeLineSyncAPIResponseModel is 环线数据同步 成功返回结果
type AlibabaAlihouseNewhomeLineSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_line_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeLineSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLineSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeLineSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLineSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeLineSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeLineSyncAPIResponse
func GetAlibabaAlihouseNewhomeLineSyncAPIResponse() *AlibabaAlihouseNewhomeLineSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeLineSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeLineSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeLineSyncAPIResponse 将 AlibabaAlihouseNewhomeLineSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeLineSyncAPIResponse(v *AlibabaAlihouseNewhomeLineSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeLineSyncAPIResponse.Put(v)
}
