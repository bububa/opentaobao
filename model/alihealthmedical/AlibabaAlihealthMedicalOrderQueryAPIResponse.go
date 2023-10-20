package alihealthmedical

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalOrderQueryAPIResponse 三方机构查询订单详情接口 API返回值
// alibaba.alihealth.medical.order.query
//
// 查询订单详情，包括评价
type AlibabaAlihealthMedicalOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMedicalOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMedicalOrderQueryAPIResponseModel).Reset()
}

// AlibabaAlihealthMedicalOrderQueryAPIResponseModel is 三方机构查询订单详情接口 成功返回结果
type AlibabaAlihealthMedicalOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMedicalOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceResult = nil
}

var poolAlibabaAlihealthMedicalOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMedicalOrderQueryAPIResponse)
	},
}

// GetAlibabaAlihealthMedicalOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMedicalOrderQueryAPIResponse
func GetAlibabaAlihealthMedicalOrderQueryAPIResponse() *AlibabaAlihealthMedicalOrderQueryAPIResponse {
	return poolAlibabaAlihealthMedicalOrderQueryAPIResponse.Get().(*AlibabaAlihealthMedicalOrderQueryAPIResponse)
}

// ReleaseAlibabaAlihealthMedicalOrderQueryAPIResponse 将 AlibabaAlihealthMedicalOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMedicalOrderQueryAPIResponse(v *AlibabaAlihealthMedicalOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMedicalOrderQueryAPIResponse.Put(v)
}
