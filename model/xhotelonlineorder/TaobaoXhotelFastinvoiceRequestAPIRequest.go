package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelFastinvoiceRequestAPIRequest
极速开票开票请求回传 API请求
taobao.xhotel.fastinvoice.request

极速开票开票请求回传,用于记录航信开票请求数据 */
type TaobaoXhotelFastinvoiceRequestAPIRequest struct {
	model.Params
	// 无
	_invoiceInfoParam *InvoiceInfoParam
}

// New
