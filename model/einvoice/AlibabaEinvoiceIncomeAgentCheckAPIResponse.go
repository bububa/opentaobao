package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeAgentCheckAPIResponse agent注册校验 API返回值
// alibaba.einvoice.income.agent.check
//
// agent注册是，需要交易用户填写的agentId是否有效
type AlibabaEinvoiceIncomeAgentCheckAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceIncomeAgentCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceIncomeAgentCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceIncomeAgentCheckAPIResponseModel).Reset()
}

// AlibabaEinvoiceIncomeAgentCheckAPIResponseModel is agent注册校验 成功返回结果
type AlibabaEinvoiceIncomeAgentCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_income_agent_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceIncomeAgentCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoiceIncomeAgentCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceIncomeAgentCheckAPIResponse)
	},
}

// GetAlibabaEinvoiceIncomeAgentCheckAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceIncomeAgentCheckAPIResponse
func GetAlibabaEinvoiceIncomeAgentCheckAPIResponse() *AlibabaEinvoiceIncomeAgentCheckAPIResponse {
	return poolAlibabaEinvoiceIncomeAgentCheckAPIResponse.Get().(*AlibabaEinvoiceIncomeAgentCheckAPIResponse)
}

// ReleaseAlibabaEinvoiceIncomeAgentCheckAPIResponse 将 AlibabaEinvoiceIncomeAgentCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceIncomeAgentCheckAPIResponse(v *AlibabaEinvoiceIncomeAgentCheckAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceIncomeAgentCheckAPIResponse.Put(v)
}
