package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest
【已废除】同步历史的销售数据 API请求
alibaba.tmallgenie.scp.plan.history.sale.qty.get

同步历史的销售数据 */
type AlibabaTmallgenieScpPlanHistorySaleQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
