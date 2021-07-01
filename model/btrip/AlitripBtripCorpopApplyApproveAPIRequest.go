package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCorpopApplyApproveAPIRequest
【商旅】更新审批单状态 API请求
alitrip.btrip.corpop.apply.approve

【商旅】更新审批单状态 */
type AlitripBtripCorpopApplyApproveAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiUpdateApplyRq
}

// New
