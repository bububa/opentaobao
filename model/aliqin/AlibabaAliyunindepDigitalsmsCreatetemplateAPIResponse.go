package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse 数字短信模板创建 API返回值
// alibaba.aliyunindep.digitalsms.createtemplate
//
// 数字短信模板创建，给阿里云一方产品使用，类型：9
type AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse struct {
	model.CommonResponse
	AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponseModel).Reset()
}

// AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponseModel is 数字短信模板创建 成功返回结果
type AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliyunindep_digitalsms_createtemplate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAliyunindepDigitalsmsCreatetemplateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse)
	},
}

// GetAlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse 从 sync.Pool 获取 AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse
func GetAlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse() *AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse {
	return poolAlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse.Get().(*AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse)
}

// ReleaseAlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse 将 AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse(v *AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse) {
	v.Reset()
	poolAlibabaAliyunindepDigitalsmsCreatetemplateAPIResponse.Put(v)
}
