package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceFlowRenewAPIResponse 工单(入驻、加盘、续约)续约 API返回值
// alibaba.einvoice.flow.renew
//
// 工单(含入驻、加盘、续约工单)续约能力开放
type AlibabaEinvoiceFlowRenewAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceFlowRenewAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowRenewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceFlowRenewAPIResponseModel).Reset()
}

// AlibabaEinvoiceFlowRenewAPIResponseModel is 工单(入驻、加盘、续约)续约 成功返回结果
type AlibabaEinvoiceFlowRenewAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_flow_renew_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowRenewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceFlowRenewAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceFlowRenewAPIResponse)
	},
}

// GetAlibabaEinvoiceFlowRenewAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceFlowRenewAPIResponse
func GetAlibabaEinvoiceFlowRenewAPIResponse() *AlibabaEinvoiceFlowRenewAPIResponse {
	return poolAlibabaEinvoiceFlowRenewAPIResponse.Get().(*AlibabaEinvoiceFlowRenewAPIResponse)
}

// ReleaseAlibabaEinvoiceFlowRenewAPIResponse 将 AlibabaEinvoiceFlowRenewAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceFlowRenewAPIResponse(v *AlibabaEinvoiceFlowRenewAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceFlowRenewAPIResponse.Put(v)
}
