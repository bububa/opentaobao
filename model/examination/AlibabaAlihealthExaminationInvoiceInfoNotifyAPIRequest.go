package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationInvoiceInfoNotifyAPIRequest
体检机构同步发票信息给阿里健康 API请求
alibaba.alihealth.examination.invoice.info.notify

体检机构向阿里健康同步发票信息 */
type AlibabaAlihealthExaminationInvoiceInfoNotifyAPIRequest struct {
	model.Params
	// 开票状态；（have_submit已提交、invoice_done已开票）
	_invoiceStatus string
	// 发票访问地址；（invoice_status在已开票状态下必填）
	_invoiceUrl string
	// 阿里健康预约凭证
	_reserveNumber string
}

// New
