package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLegalCaseStandpointFeedbackAPIRequest
新增或更新 反馈口径(采纳口径/不采纳口径) API请求
alibaba.legal.case.standpoint.feedback

新增或更新 反馈口径(采纳口径/不采纳口径) */
type AlibabaLegalCaseStandpointFeedbackAPIRequest struct {
	model.Params
	// 反馈对象
	_feedbackRequestModel *FeedbackRequestModel
}

// New
