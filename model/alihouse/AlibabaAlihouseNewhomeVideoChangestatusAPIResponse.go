package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeVideoChangestatusAPIResponse 视频草稿状态更新 API返回值
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
type AlibabaAlihouseNewhomeVideoChangestatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeVideoChangestatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeVideoChangestatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeVideoChangestatusAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeVideoChangestatusAPIResponseModel is 视频草稿状态更新 成功返回结果
type AlibabaAlihouseNewhomeVideoChangestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_video_changestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeVideoChangestatusResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeVideoChangestatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeVideoChangestatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeVideoChangestatusAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeVideoChangestatusAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeVideoChangestatusAPIResponse
func GetAlibabaAlihouseNewhomeVideoChangestatusAPIResponse() *AlibabaAlihouseNewhomeVideoChangestatusAPIResponse {
	return poolAlibabaAlihouseNewhomeVideoChangestatusAPIResponse.Get().(*AlibabaAlihouseNewhomeVideoChangestatusAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeVideoChangestatusAPIResponse 将 AlibabaAlihouseNewhomeVideoChangestatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeVideoChangestatusAPIResponse(v *AlibabaAlihouseNewhomeVideoChangestatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeVideoChangestatusAPIResponse.Put(v)
}
