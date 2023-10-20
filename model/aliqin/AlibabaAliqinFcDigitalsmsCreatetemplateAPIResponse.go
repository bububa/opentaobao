package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfcdigitalsmscreatetemplateAPIResponse 数字短信模板创建 API返回值
// alibaba.aliqin.fc.digitalsms.createtemplate
//
// 数字短信模板创建
type AlibabaaliqinfcdigitalsmscreatetemplateAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinfcdigitalsmscreatetemplateAPIResponseModel
}

// AlibabaaliqinfcdigitalsmscreatetemplateAPIResponseModel is 数字短信模板创建 成功返回结果
type AlibabaaliqinfcdigitalsmscreatetemplateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_digitalsms_createtemplate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlibabaaliqinfcdigitalsmscreatetemplateResult `json:"result,omitempty" xml:"result,omitempty"`
}
