package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLayoutSyncAPIResponse 房通户型数据同步 API返回值
// alibaba.alihouse.newhome.layout.sync
//
// 房通户型数据同步
type AlibabaAlihouseNewhomeLayoutSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeLayoutSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLayoutSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeLayoutSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeLayoutSyncAPIResponseModel is 房通户型数据同步 成功返回结果
type AlibabaAlihouseNewhomeLayoutSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_layout_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeLayoutSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLayoutSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeLayoutSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLayoutSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeLayoutSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeLayoutSyncAPIResponse
func GetAlibabaAlihouseNewhomeLayoutSyncAPIResponse() *AlibabaAlihouseNewhomeLayoutSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeLayoutSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeLayoutSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeLayoutSyncAPIResponse 将 AlibabaAlihouseNewhomeLayoutSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeLayoutSyncAPIResponse(v *AlibabaAlihouseNewhomeLayoutSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeLayoutSyncAPIResponse.Put(v)
}
