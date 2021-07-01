package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest
定向推广-新增或删除屏蔽词 API请求
alibaba.scbp.target.ad.plan.forbidden.word.modify

定向推广-新增或删除屏蔽词 */
type AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest struct {
	model.Params
	// TopP4pQuickForbiddenWord
	_topP4pQuickForbiddenWord *TopP4pQuickForbiddenWordDto
}

// New
