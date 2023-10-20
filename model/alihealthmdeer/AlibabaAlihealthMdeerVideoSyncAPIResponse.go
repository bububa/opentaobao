package alihealthmdeer

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMdeerVideoSyncAPIResponse 合作伙伴视频同步给医知鹿接口 API返回值
// alibaba.alihealth.mdeer.video.sync
//
// 合伙做伴内容同步接口，用来视频同步
type AlibabaAlihealthMdeerVideoSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMdeerVideoSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMdeerVideoSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMdeerVideoSyncAPIResponseModel).Reset()
}

// AlibabaAlihealthMdeerVideoSyncAPIResponseModel is 合作伙伴视频同步给医知鹿接口 成功返回结果
type AlibabaAlihealthMdeerVideoSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_mdeer_video_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMdeerVideoSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMdeerVideoSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMdeerVideoSyncAPIResponse)
	},
}

// GetAlibabaAlihealthMdeerVideoSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMdeerVideoSyncAPIResponse
func GetAlibabaAlihealthMdeerVideoSyncAPIResponse() *AlibabaAlihealthMdeerVideoSyncAPIResponse {
	return poolAlibabaAlihealthMdeerVideoSyncAPIResponse.Get().(*AlibabaAlihealthMdeerVideoSyncAPIResponse)
}

// ReleaseAlibabaAlihealthMdeerVideoSyncAPIResponse 将 AlibabaAlihealthMdeerVideoSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMdeerVideoSyncAPIResponse(v *AlibabaAlihealthMdeerVideoSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMdeerVideoSyncAPIResponse.Put(v)
}
