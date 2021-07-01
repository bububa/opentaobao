package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusInvoiceReturnAPIRequest
发票回调接口 API请求
taobao.bus.invoice.return

汽车票发票回调接口 */
type TaobaoBusInvoiceReturnAPIRequest struct {
	model.Params
	// 入参对象
	_invoiceParam *ReceiptDo
}

// New
