package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeBusinessSyncAPIResponse 商圈数据同步 API返回值
// alibaba.alihouse.newhome.business.sync
//
// 商圈数据同步
type AlibabaAlihouseNewhomeBusinessSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeBusinessSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel is 商圈数据同步 成功返回结果
type AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_business_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeBusinessSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeBusinessSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeBusinessSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeBusinessSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeBusinessSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeBusinessSyncAPIResponse
func GetAlibabaAlihouseNewhomeBusinessSyncAPIResponse() *AlibabaAlihouseNewhomeBusinessSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeBusinessSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeBusinessSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeBusinessSyncAPIResponse 将 AlibabaAlihouseNewhomeBusinessSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeBusinessSyncAPIResponse(v *AlibabaAlihouseNewhomeBusinessSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeBusinessSyncAPIResponse.Put(v)
}
