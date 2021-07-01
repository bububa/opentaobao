package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceInvoiceapplyUpdateAPIRequest
商家开票申请单状态回传 API请求
alibaba.einvoice.invoiceapply.update

开票服务商更新商家开票申请单状态 */
type AlibabaEinvoiceInvoiceapplyUpdateAPIRequest struct {
	model.Params
	// 申请单id
	_applyId string
	// 申请单状态：1：待确认，2：开票中，3：拒绝开票，4：发票已发出，0：完成开票
	_status int64
	// 扩展信息,目前用于回传文本及物流消息
	_exInfo string
}

// New
