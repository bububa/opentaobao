package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增或更新 反馈口径(采纳口径/不采纳口径) API请求
alibaba.legal.case.standpoint.feedback

新增或更新 反馈口径(采纳口径/不采纳口径)
*/
type AlibabaLegalCaseStandpointFeedbackRequest struct {
    model.Params
    // 反馈对象
    _feedbackRequestModel   *FeedbackRequestModel
}

// 初始化AlibabaLegalCaseStandpointFeedbackRequest对象
func NewAlibabaLegalCaseStandpointFeedbackRequest() *AlibabaLegalCaseStandpointFeedbackRequest{
    return &AlibabaLegalCaseStandpointFeedbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseStandpointFeedbackRequest) GetApiMethodName() string {
    return "alibaba.legal.case.standpoint.feedback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseStandpointFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FeedbackRequestModel Setter
// 反馈对象
func (r *AlibabaLegalCaseStandpointFeedbackRequest) SetFeedbackRequestModel(_feedbackRequestModel *FeedbackRequestModel) error {
    r._feedbackRequestModel = _feedbackRequestModel
    r.Set("feedback_request_model", _feedbackRequestModel)
    return nil
}

// FeedbackRequestModel Getter
func (r AlibabaLegalCaseStandpointFeedbackRequest) GetFeedbackRequestModel() *FeedbackRequestModel {
    return r._feedbackRequestModel
}
