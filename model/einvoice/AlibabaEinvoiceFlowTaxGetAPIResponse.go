package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceFlowTaxGetAPIResponse 查询税控开通工单详情 API返回值
// alibaba.einvoice.flow.tax.get
//
// 查询税控开通工单详情，接口返回工单状态、开票商户信息以及税控设备信息。
// 场景使用：1、业务前台收到入驻成功消息后，调用此接口查询最终的商户信息和设备信息；2、主动补偿查询：当工单长时间未收到事件通知，可能存在丢消息的情况，此时可主动查询该工单，更新本地工单状态。
type AlibabaEinvoiceFlowTaxGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceFlowTaxGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowTaxGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceFlowTaxGetAPIResponseModel).Reset()
}

// AlibabaEinvoiceFlowTaxGetAPIResponseModel is 查询税控开通工单详情 成功返回结果
type AlibabaEinvoiceFlowTaxGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_flow_tax_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceFlowTaxGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceFlowTaxGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceFlowTaxGetAPIResponse)
	},
}

// GetAlibabaEinvoiceFlowTaxGetAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceFlowTaxGetAPIResponse
func GetAlibabaEinvoiceFlowTaxGetAPIResponse() *AlibabaEinvoiceFlowTaxGetAPIResponse {
	return poolAlibabaEinvoiceFlowTaxGetAPIResponse.Get().(*AlibabaEinvoiceFlowTaxGetAPIResponse)
}

// ReleaseAlibabaEinvoiceFlowTaxGetAPIResponse 将 AlibabaEinvoiceFlowTaxGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceFlowTaxGetAPIResponse(v *AlibabaEinvoiceFlowTaxGetAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceFlowTaxGetAPIResponse.Put(v)
}
