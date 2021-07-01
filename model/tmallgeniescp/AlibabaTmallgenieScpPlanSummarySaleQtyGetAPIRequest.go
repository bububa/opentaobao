package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest
同步销售数据按照渠道类型汇总 API请求
alibaba.tmallgenie.scp.plan.summary.sale.qty.get

同步销售数据按照渠道类型汇总 */
type AlibabaTmallgenieScpPlanSummarySaleQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
