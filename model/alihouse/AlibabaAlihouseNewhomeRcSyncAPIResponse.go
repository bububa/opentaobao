package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRcSyncAPIResponse 阿里房产图文草稿信息同步 API返回值
// alibaba.alihouse.newhome.rc.sync
//
// 接收图文草稿信息
type AlibabaAlihouseNewhomeRcSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeRcSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeRcSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeRcSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeRcSyncAPIResponseModel is 阿里房产图文草稿信息同步 成功返回结果
type AlibabaAlihouseNewhomeRcSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_rc_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeRcSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeRcSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeRcSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRcSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeRcSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeRcSyncAPIResponse
func GetAlibabaAlihouseNewhomeRcSyncAPIResponse() *AlibabaAlihouseNewhomeRcSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeRcSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeRcSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeRcSyncAPIResponse 将 AlibabaAlihouseNewhomeRcSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRcSyncAPIResponse(v *AlibabaAlihouseNewhomeRcSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRcSyncAPIResponse.Put(v)
}
