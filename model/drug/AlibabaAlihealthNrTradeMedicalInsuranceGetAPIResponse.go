package drug

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse 阿里健康医保支付信息获取 API返回值
// alibaba.alihealth.nr.trade.medical.insurance.get
//
// 阿里健康医保支付信息获取
type AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponseModel).Reset()
}

// AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponseModel is 阿里健康医保支付信息获取 成功返回结果
type AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_trade_medical_insurance_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值总
	Result *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse)
	},
}

// GetAlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse
func GetAlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse() *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse {
	return poolAlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse.Get().(*AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse)
}

// ReleaseAlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse 将 AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse(v *AlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthNrTradeMedicalInsuranceGetAPIResponse.Put(v)
}
