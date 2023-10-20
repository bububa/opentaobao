package ascm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse 英迈发票同步到结算 API返回值
// alibaba.ascm.settlement.invoice.synchronization.im
//
// 外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
type AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse struct {
	model.CommonResponse
	AlibabaAscmSettlementInvoiceSynchronizationImAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscmSettlementInvoiceSynchronizationImAPIResponseModel).Reset()
}

// AlibabaAscmSettlementInvoiceSynchronizationImAPIResponseModel is 英迈发票同步到结算 成功返回结果
type AlibabaAscmSettlementInvoiceSynchronizationImAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascm_settlement_invoice_synchronization_im_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SettlementGatewayResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscmSettlementInvoiceSynchronizationImAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscmSettlementInvoiceSynchronizationImAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse)
	},
}

// GetAlibabaAscmSettlementInvoiceSynchronizationImAPIResponse 从 sync.Pool 获取 AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse
func GetAlibabaAscmSettlementInvoiceSynchronizationImAPIResponse() *AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse {
	return poolAlibabaAscmSettlementInvoiceSynchronizationImAPIResponse.Get().(*AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse)
}

// ReleaseAlibabaAscmSettlementInvoiceSynchronizationImAPIResponse 将 AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscmSettlementInvoiceSynchronizationImAPIResponse(v *AlibabaAscmSettlementInvoiceSynchronizationImAPIResponse) {
	v.Reset()
	poolAlibabaAscmSettlementInvoiceSynchronizationImAPIResponse.Put(v)
}
