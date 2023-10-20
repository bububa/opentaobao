package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse 数字短信模板创建 API返回值
// alibaba.aliqin.fc.digitalsms.createtemplate
//
// 数字短信模板创建
type AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponseModel).Reset()
}

// AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponseModel is 数字短信模板创建 成功返回结果
type AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_digitalsms_createtemplate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAliqinFcDigitalsmsCreatetemplateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse)
	},
}

// GetAlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse
func GetAlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse() *AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse {
	return poolAlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse.Get().(*AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse)
}

// ReleaseAlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse 将 AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse(v *AlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcDigitalsmsCreatetemplateAPIResponse.Put(v)
}
