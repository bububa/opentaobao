package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTemplateQuesGetAPIResponse 获取SPU最新版本问卷 API返回值
// alibaba.idle.template.ques.get
//
// 获取SPU最新版本问卷
type AlibabaIdleTemplateQuesGetAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTemplateQuesGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTemplateQuesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTemplateQuesGetAPIResponseModel).Reset()
}

// AlibabaIdleTemplateQuesGetAPIResponseModel is 获取SPU最新版本问卷 成功返回结果
type AlibabaIdleTemplateQuesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_template_ques_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTemplateQuesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTemplateQuesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTemplateQuesGetAPIResponse)
	},
}

// GetAlibabaIdleTemplateQuesGetAPIResponse 从 sync.Pool 获取 AlibabaIdleTemplateQuesGetAPIResponse
func GetAlibabaIdleTemplateQuesGetAPIResponse() *AlibabaIdleTemplateQuesGetAPIResponse {
	return poolAlibabaIdleTemplateQuesGetAPIResponse.Get().(*AlibabaIdleTemplateQuesGetAPIResponse)
}

// ReleaseAlibabaIdleTemplateQuesGetAPIResponse 将 AlibabaIdleTemplateQuesGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTemplateQuesGetAPIResponse(v *AlibabaIdleTemplateQuesGetAPIResponse) {
	v.Reset()
	poolAlibabaIdleTemplateQuesGetAPIResponse.Put(v)
}
