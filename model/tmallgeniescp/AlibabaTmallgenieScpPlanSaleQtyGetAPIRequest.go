package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest
12-同步销售数据 API请求
alibaba.tmallgenie.scp.plan.sale.qty.get

同步销售数据 */
type AlibabaTmallgenieScpPlanSaleQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
