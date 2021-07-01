package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest
11-同步本周的po单（从W-1周到W+4周） API请求
alibaba.tmallgenie.scp.plan.current.po.get

11-同步本周的po单（从W-1周到W+4周） */
type AlibabaTmallgenieScpPlanCurrentPoGetAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
