package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFcMallxInteractionAiPicListAPIResponse 花园ai作画定制查询 API返回值
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
type AlibabaFcMallxInteractionAiPicListAPIResponse struct {
	model.CommonResponse
	AlibabaFcMallxInteractionAiPicListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFcMallxInteractionAiPicListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFcMallxInteractionAiPicListAPIResponseModel).Reset()
}

// AlibabaFcMallxInteractionAiPicListAPIResponseModel is 花园ai作画定制查询 成功返回结果
type AlibabaFcMallxInteractionAiPicListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fc_mallx_interaction_ai_pic_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaFcMallxInteractionAiPicListResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFcMallxInteractionAiPicListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFcMallxInteractionAiPicListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFcMallxInteractionAiPicListAPIResponse)
	},
}

// GetAlibabaFcMallxInteractionAiPicListAPIResponse 从 sync.Pool 获取 AlibabaFcMallxInteractionAiPicListAPIResponse
func GetAlibabaFcMallxInteractionAiPicListAPIResponse() *AlibabaFcMallxInteractionAiPicListAPIResponse {
	return poolAlibabaFcMallxInteractionAiPicListAPIResponse.Get().(*AlibabaFcMallxInteractionAiPicListAPIResponse)
}

// ReleaseAlibabaFcMallxInteractionAiPicListAPIResponse 将 AlibabaFcMallxInteractionAiPicListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFcMallxInteractionAiPicListAPIResponse(v *AlibabaFcMallxInteractionAiPicListAPIResponse) {
	v.Reset()
	poolAlibabaFcMallxInteractionAiPicListAPIResponse.Put(v)
}
