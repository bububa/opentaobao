package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest
二级物料-LT内的POGAP数据回传 API请求
alibaba.tmallgenie.scp.plan.rawpo.gap.return

二级物料-LT内的POGAP数据回传 */
type AlibabaTmallgenieScpPlanRawpoGapReturnAPIRequest struct {
	model.Params
	// 请求对象
	_rawPogapRequest *RawPurchaseOrderGapRequest
}

// New
