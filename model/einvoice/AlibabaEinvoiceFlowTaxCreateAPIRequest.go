package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceFlowTaxCreateAPIRequest
创建税控开通工单 API请求
alibaba.einvoice.flow.tax.create

商户在业务前台订购税控产品后，调用阿里发票此接口，提交税号的入驻开通工单。此接口返回为工单的提交结果，非真正入驻结果。开通结果会在商户完成设备的部署安装 入驻完成后，由阿里发票通过消息异步通知到业务前台。 */
type AlibabaEinvoiceFlowTaxCreateAPIRequest struct {
	model.Params
	// 工单请求
	_invoiceTaxFlowCreateDto *InvoiceTaxFlowCreateDto
}

// New
