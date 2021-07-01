package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest
【已废除】11-同步历史所有的po单 API请求
alibaba.tmallgenie.scp.plan.history.po.get

同步历史po单 */
type AlibabaTmallgenieScpPlanHistoryPoGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
