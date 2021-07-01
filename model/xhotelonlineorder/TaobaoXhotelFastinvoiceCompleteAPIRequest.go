package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelFastinvoiceCompleteAPIRequest
极速开票开票请求完成 API请求
taobao.xhotel.fastinvoice.complete

极速开票开票请求回传,用于更新航信开票请求数据 */
type TaobaoXhotelFastinvoiceCompleteAPIRequest struct {
	model.Params
	// 无
	_invoiceInfoParam *InvoiceInfoParam
}

// New
