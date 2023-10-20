package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseQuerystandpointSaveAPIRequest 法宝侧主动查询反馈 API请求
// alibaba.legal.case.querystandpoint.save
//
// 法宝侧主动查询反馈口径,此接口只用来反馈主动查询的口径,之前推送的口径反馈不适合
type AlibabaLegalCaseQuerystandpointSaveAPIRequest struct {
	model.Params
	// 反馈列表
	_feedbackRequestModels []FeedbackRequestModel
}

// NewAlibabaLegalCaseQuerystandpointSaveRequest 初始化AlibabaLegalCaseQuerystandpointSaveAPIRequest对象
func NewAlibabaLegalCaseQuerystandpointSaveRequest() *AlibabaLegalCaseQuerystandpointSaveAPIRequest {
	return &AlibabaLegalCaseQuerystandpointSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseQuerystandpointSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.querystandpoint.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseQuerystandpointSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseQuerystandpointSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeedbackRequestModels is FeedbackRequestModels Setter
// 反馈列表
func (r *AlibabaLegalCaseQuerystandpointSaveAPIRequest) SetFeedbackRequestModels(_feedbackRequestModels []FeedbackRequestModel) error {
	r._feedbackRequestModels = _feedbackRequestModels
	r.Set("feedback_request_models", _feedbackRequestModels)
	return nil
}

// GetFeedbackRequestModels FeedbackRequestModels Getter
func (r AlibabaLegalCaseQuerystandpointSaveAPIRequest) GetFeedbackRequestModels() []FeedbackRequestModel {
	return r._feedbackRequestModels
}
