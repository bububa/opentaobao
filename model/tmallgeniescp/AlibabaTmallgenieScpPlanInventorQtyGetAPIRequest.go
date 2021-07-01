package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest
10-同步库存现有量 API请求
alibaba.tmallgenie.scp.plan.inventor.qty.get

同步库存现有量 */
type AlibabaTmallgenieScpPlanInventorQtyGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
