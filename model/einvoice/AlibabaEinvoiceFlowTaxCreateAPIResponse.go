package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceFlowTaxCreateAPIResponse 创建税控开通工单 API返回值
// alibaba.einvoice.flow.tax.create
//
// 商户在业务前台订购税控产品后，调用阿里发票此接口，提交税号的入驻开通工单。此接口返回为工单的提交结果，非真正入驻结果。开通结果会在商户完成设备的部署安装 入驻完成后，由阿里发票通过消息异步通知到业务前台。
type AlibabaEinvoiceFlowTaxCreateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceFlowTaxCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowTaxCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceFlowTaxCreateAPIResponseModel).Reset()
}

// AlibabaEinvoiceFlowTaxCreateAPIResponseModel is 创建税控开通工单 成功返回结果
type AlibabaEinvoiceFlowTaxCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_flow_tax_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 工单ID，发票中台生成
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowTaxCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FlowId = ""
}

var poolAlibabaEinvoiceFlowTaxCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceFlowTaxCreateAPIResponse)
	},
}

// GetAlibabaEinvoiceFlowTaxCreateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceFlowTaxCreateAPIResponse
func GetAlibabaEinvoiceFlowTaxCreateAPIResponse() *AlibabaEinvoiceFlowTaxCreateAPIResponse {
	return poolAlibabaEinvoiceFlowTaxCreateAPIResponse.Get().(*AlibabaEinvoiceFlowTaxCreateAPIResponse)
}

// ReleaseAlibabaEinvoiceFlowTaxCreateAPIResponse 将 AlibabaEinvoiceFlowTaxCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceFlowTaxCreateAPIResponse(v *AlibabaEinvoiceFlowTaxCreateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceFlowTaxCreateAPIResponse.Put(v)
}
