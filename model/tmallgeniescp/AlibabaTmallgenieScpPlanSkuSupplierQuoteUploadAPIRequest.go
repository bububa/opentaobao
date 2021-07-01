package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest
标准供应商配额同步 API请求
alibaba.tmallgenie.scp.plan.sku.supplier.quote.upload

标准供应商配额同步 */
type AlibabaTmallgenieScpPlanSkuSupplierQuoteUploadAPIRequest struct {
	model.Params
	// 对象
	_netDemandRequest *NetDemandRequest
}

// New
