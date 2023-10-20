package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTemplateQuesOnlineAPIResponse 预上线SPU问卷版本 API返回值
// alibaba.idle.template.ques.online
//
// 获取SPU最新版本问卷
type AlibabaIdleTemplateQuesOnlineAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTemplateQuesOnlineAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTemplateQuesOnlineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTemplateQuesOnlineAPIResponseModel).Reset()
}

// AlibabaIdleTemplateQuesOnlineAPIResponseModel is 预上线SPU问卷版本 成功返回结果
type AlibabaIdleTemplateQuesOnlineAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_template_ques_online_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTemplateQuesOnlineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTemplateQuesOnlineAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTemplateQuesOnlineAPIResponse)
	},
}

// GetAlibabaIdleTemplateQuesOnlineAPIResponse 从 sync.Pool 获取 AlibabaIdleTemplateQuesOnlineAPIResponse
func GetAlibabaIdleTemplateQuesOnlineAPIResponse() *AlibabaIdleTemplateQuesOnlineAPIResponse {
	return poolAlibabaIdleTemplateQuesOnlineAPIResponse.Get().(*AlibabaIdleTemplateQuesOnlineAPIResponse)
}

// ReleaseAlibabaIdleTemplateQuesOnlineAPIResponse 将 AlibabaIdleTemplateQuesOnlineAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTemplateQuesOnlineAPIResponse(v *AlibabaIdleTemplateQuesOnlineAPIResponse) {
	v.Reset()
	poolAlibabaIdleTemplateQuesOnlineAPIResponse.Put(v)
}
