package alihealthmedical

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalOrderRefuseAPIResponse 三方机构通知平台"医生拒诊" API返回值
// alibaba.alihealth.medical.order.refuse
//
// 三方机构通知平台&#34;医生拒诊&#34;
type AlibabaAlihealthMedicalOrderRefuseAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalOrderRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalOrderRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalOrderRefuseAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalOrderRefuseAPIResponseModel is 三方机构通知平台"医生拒诊" 成功返回结果
type AlibabaAlihealthMedicalOrderRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_order_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalOrderRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthMedicalOrderRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalOrderRefuseAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalOrderRefuseAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalOrderRefuseAPIResponse
func GetAlibabaAlihealthMedicalOrderRefuseAPIResponse() *AlibabaAlihealthMedicalOrderRefuseAPIResponse {
	return poolAlibabaAlihealthMedicalOrderRefuseAPIResponse.Get().(*AlibabaAlihealthMedicalOrderRefuseAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalOrderRefuseAPIResponse 将 AlibabaAlihealthMedicalOrderRefuseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalOrderRefuseAPIResponse(v *AlibabaAlihealthMedicalOrderRefuseAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalOrderRefuseAPIResponse.Put(v)
}
