package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest
供应商校准后的配额同步 API请求
alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload

供应商校准后的配额同步 */
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// New
