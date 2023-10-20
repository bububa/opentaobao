package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasestandpointfeedbackAPIRequest 新增或更新 反馈口径(采纳口径/不采纳口径) API请求
// alibaba.legal.case.standpoint.feedback
//
// 新增或更新 反馈口径(采纳口径/不采纳口径)
type AlibabalegalcasestandpointfeedbackAPIRequest struct {
	model.Params
	// 反馈对象
	_feedbackRequestModel *FeedbackRequestModel
}

// NewAlibabalegalcasestandpointfeedbackRequest 初始化AlibabalegalcasestandpointfeedbackAPIRequest对象
func NewAlibabalegalcasestandpointfeedbackRequest() *AlibabalegalcasestandpointfeedbackAPIRequest {
	return &AlibabalegalcasestandpointfeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcasestandpointfeedbackAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.standpoint.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcasestandpointfeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcasestandpointfeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeedbackRequestModel is FeedbackRequestModel Setter
// 反馈对象
func (r *AlibabalegalcasestandpointfeedbackAPIRequest) SetFeedbackRequestModel(_feedbackRequestModel *FeedbackRequestModel) error {
	r._feedbackRequestModel = _feedbackRequestModel
	r.Set("feedback_request_model", _feedbackRequestModel)
	return nil
}

// GetFeedbackRequestModel FeedbackRequestModel Getter
func (r AlibabalegalcasestandpointfeedbackAPIRequest) GetFeedbackRequestModel() *FeedbackRequestModel {
	return r._feedbackRequestModel
}
