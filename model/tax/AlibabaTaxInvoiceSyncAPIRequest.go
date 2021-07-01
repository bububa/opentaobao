package tax

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTaxInvoiceSyncAPIRequest
第三方开票回调API API请求
alibaba.tax.invoice.sync

该接口只提供给俄罗斯供应商开具发票使用，请勿申请。 */
type AlibabaTaxInvoiceSyncAPIRequest struct {
	model.Params
	// 回调发票信息，请求参数详情见链接：https://yuque.antfin-inc.com/tax/biw99l/uestb7
	_paramJson string
}

// New
