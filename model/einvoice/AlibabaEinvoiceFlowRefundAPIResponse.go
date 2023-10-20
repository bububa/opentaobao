package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceFlowRefundAPIResponse 退订工单(入驻、加盘、续约) API返回值
// alibaba.einvoice.flow.refund
//
// 电子发票工单系统，工单退订能力开放
type AlibabaEinvoiceFlowRefundAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceFlowRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceFlowRefundAPIResponseModel).Reset()
}

// AlibabaEinvoiceFlowRefundAPIResponseModel is 退订工单(入驻、加盘、续约) 成功返回结果
type AlibabaEinvoiceFlowRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_flow_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceFlowRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceFlowRefundAPIResponse)
	},
}

// GetAlibabaEinvoiceFlowRefundAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceFlowRefundAPIResponse
func GetAlibabaEinvoiceFlowRefundAPIResponse() *AlibabaEinvoiceFlowRefundAPIResponse {
	return poolAlibabaEinvoiceFlowRefundAPIResponse.Get().(*AlibabaEinvoiceFlowRefundAPIResponse)
}

// ReleaseAlibabaEinvoiceFlowRefundAPIResponse 将 AlibabaEinvoiceFlowRefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceFlowRefundAPIResponse(v *AlibabaEinvoiceFlowRefundAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceFlowRefundAPIResponse.Put(v)
}
