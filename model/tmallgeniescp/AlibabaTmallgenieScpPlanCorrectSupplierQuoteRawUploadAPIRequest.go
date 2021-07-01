package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest
同步供应商校准后的配额-二级物料 API请求
alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload

同步供应商校准后的配额-二级物料 */
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIRequest struct {
	model.Params
	// 对象
	_currentQuoteRawRequest *AbstractRequest
}

// New
