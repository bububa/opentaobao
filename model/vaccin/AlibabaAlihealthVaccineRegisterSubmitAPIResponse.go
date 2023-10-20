package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineRegisterSubmitAPIResponse cdc回传疫苗登记数据 API返回值
// alibaba.alihealth.vaccine.register.submit
//
// cdc回传疫苗登记信息
type AlibabaAlihealthVaccineRegisterSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthVaccineRegisterSubmitAPIResponseModel
}

// AlibabaAlihealthVaccineRegisterSubmitAPIResponseModel is cdc回传疫苗登记数据 成功返回结果
type AlibabaAlihealthVaccineRegisterSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_vaccine_register_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无
	Result *AlibabaAlihealthVaccineRegisterSubmitMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
