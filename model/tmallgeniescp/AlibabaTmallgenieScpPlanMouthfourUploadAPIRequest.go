package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest
21-M+4PR 回传接口接口 API请求
alibaba.tmallgenie.scp.plan.mouthfour.upload

M+4 PR 回传接口 */
type AlibabaTmallgenieScpPlanMouthfourUploadAPIRequest struct {
	model.Params
	// 请求参数
	_monthFourPrRequest *MonthFourPrRequest
}

// New
