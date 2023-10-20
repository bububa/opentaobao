package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractCurrentMixusernickAPIResponse 手淘混淆nick开放接口鉴权专用 API返回值
// alibaba.interact.current.mixusernick
//
// 手淘混淆nick开放接口鉴权专用，无数据输入输出。
type AlibabaInteractCurrentMixusernickAPIResponse struct {
	model.CommonResponse
	AlibabaInteractCurrentMixusernickAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractCurrentMixusernickAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractCurrentMixusernickAPIResponseModel).Reset()
}

// AlibabaInteractCurrentMixusernickAPIResponseModel is 手淘混淆nick开放接口鉴权专用 成功返回结果
type AlibabaInteractCurrentMixusernickAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_current_mixusernick_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractCurrentMixusernickAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractCurrentMixusernickAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractCurrentMixusernickAPIResponse)
	},
}

// GetAlibabaInteractCurrentMixusernickAPIResponse 从 sync.Pool 获取 AlibabaInteractCurrentMixusernickAPIResponse
func GetAlibabaInteractCurrentMixusernickAPIResponse() *AlibabaInteractCurrentMixusernickAPIResponse {
	return poolAlibabaInteractCurrentMixusernickAPIResponse.Get().(*AlibabaInteractCurrentMixusernickAPIResponse)
}

// ReleaseAlibabaInteractCurrentMixusernickAPIResponse 将 AlibabaInteractCurrentMixusernickAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractCurrentMixusernickAPIResponse(v *AlibabaInteractCurrentMixusernickAPIResponse) {
	v.Reset()
	poolAlibabaInteractCurrentMixusernickAPIResponse.Put(v)
}
