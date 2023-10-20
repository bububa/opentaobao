package ascm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascmsettlementinvoicesynchronizationimAPIResponse 英迈发票同步到结算 API返回值
// alibaba.ascm.settlement.invoice.synchronization.im
//
// 外部供应商通过此API将发货的发票信息推送给供应链中台结算系统
type AlibabaascmsettlementinvoicesynchronizationimAPIResponse struct {
	model.CommonResponse
	AlibabaascmsettlementinvoicesynchronizationimAPIResponseModel
}

// AlibabaascmsettlementinvoicesynchronizationimAPIResponseModel is 英迈发票同步到结算 成功返回结果
type AlibabaascmsettlementinvoicesynchronizationimAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascm_settlement_invoice_synchronization_im_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SettlementGatewayResult `json:"result,omitempty" xml:"result,omitempty"`
}
