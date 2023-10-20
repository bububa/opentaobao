package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIsvDigitalsmsCreatetemplateAPIResponse 数字短信模板创建 API返回值
// alibaba.isv.digitalsms.createtemplate
//
// 数字短信模板创建，给聚石塔，类型：2
type AlibabaIsvDigitalsmsCreatetemplateAPIResponse struct {
	model.CommonResponse
	AlibabaIsvDigitalsmsCreatetemplateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIsvDigitalsmsCreatetemplateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIsvDigitalsmsCreatetemplateAPIResponseModel).Reset()
}

// AlibabaIsvDigitalsmsCreatetemplateAPIResponseModel is 数字短信模板创建 成功返回结果
type AlibabaIsvDigitalsmsCreatetemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_isv_digitalsms_createtemplate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaIsvDigitalsmsCreatetemplateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIsvDigitalsmsCreatetemplateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIsvDigitalsmsCreatetemplateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIsvDigitalsmsCreatetemplateAPIResponse)
	},
}

// GetAlibabaIsvDigitalsmsCreatetemplateAPIResponse 从 sync.Pool 获取 AlibabaIsvDigitalsmsCreatetemplateAPIResponse
func GetAlibabaIsvDigitalsmsCreatetemplateAPIResponse() *AlibabaIsvDigitalsmsCreatetemplateAPIResponse {
	return poolAlibabaIsvDigitalsmsCreatetemplateAPIResponse.Get().(*AlibabaIsvDigitalsmsCreatetemplateAPIResponse)
}

// ReleaseAlibabaIsvDigitalsmsCreatetemplateAPIResponse 将 AlibabaIsvDigitalsmsCreatetemplateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIsvDigitalsmsCreatetemplateAPIResponse(v *AlibabaIsvDigitalsmsCreatetemplateAPIResponse) {
	v.Reset()
	poolAlibabaIsvDigitalsmsCreatetemplateAPIResponse.Put(v)
}
