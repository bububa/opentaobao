package ascm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest
英迈发票同步到结算 API请求
alibaba.ascm.settlement.invoice.synchronization.im

外部供应商通过此API将发货的发票信息推送给供应链中台结算系统 */
type AlibabaAscmSettlementInvoiceSynchronizationImAPIRequest struct {
	model.Params
	// im invoice xml
	_xmlDataSlot string
}

// New
