package aliqin

import (
	"encoding/xml"

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

// AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponseModel is 数字短信模板创建 成功返回结果
type AlibabaAliyunindepDigitalsmsCreatetemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliyunindep_digitalsms_createtemplate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaAliyunindepDigitalsmsCreatetemplateResult `json:"result,omitempty" xml:"result,omitempty"`
}
