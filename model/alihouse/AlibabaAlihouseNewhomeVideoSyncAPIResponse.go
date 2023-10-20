package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeVideoSyncAPIResponse 视频草稿信息同步 API返回值
// alibaba.alihouse.newhome.video.sync
//
// 接收视频信息记录
type AlibabaAlihouseNewhomeVideoSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeVideoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeVideoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeVideoSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeVideoSyncAPIResponseModel is 视频草稿信息同步 成功返回结果
type AlibabaAlihouseNewhomeVideoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_video_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeVideoSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeVideoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeVideoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeVideoSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeVideoSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeVideoSyncAPIResponse
func GetAlibabaAlihouseNewhomeVideoSyncAPIResponse() *AlibabaAlihouseNewhomeVideoSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeVideoSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeVideoSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeVideoSyncAPIResponse 将 AlibabaAlihouseNewhomeVideoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeVideoSyncAPIResponse(v *AlibabaAlihouseNewhomeVideoSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeVideoSyncAPIResponse.Put(v)
}
