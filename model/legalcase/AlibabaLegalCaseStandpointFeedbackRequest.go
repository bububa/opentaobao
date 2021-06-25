package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增或更新 反馈口径(采纳口径/不采纳口径) APIRequest
alibaba.legal.case.standpoint.feedback

新增或更新 反馈口径(采纳口径/不采纳口径)
*/
type AlibabaLegalCaseStandpointFeedbackRequest struct {
    model.Params

    // 反馈对象
    feedbackRequestModel   *FeedbackRequestModel 

}

func NewAlibabaLegalCaseStandpointFeedbackRequest() *AlibabaLegalCaseStandpointFeedbackRequest{
    return &AlibabaLegalCaseStandpointFeedbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalCaseStandpointFeedbackRequest) GetApiMethodName() string {
    return "alibaba.legal.case.standpoint.feedback"
}

func (r AlibabaLegalCaseStandpointFeedbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalCaseStandpointFeedbackRequest) SetFeedbackRequestModel(feedbackRequestModel *FeedbackRequestModel) error {
    r.feedbackRequestModel = feedbackRequestModel
    r.Set("feedback_request_model", feedbackRequestModel)
    return nil
}

func (r AlibabaLegalCaseStandpointFeedbackRequest) GetFeedbackRequestModel() *FeedbackRequestModel {
    return r.feedbackRequestModel
}

