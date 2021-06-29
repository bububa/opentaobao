package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫好房楼盘评测同步 APIRequest
alibaba.alihouse.newhome.review.sync

接受楼盘测评图文信息
*/
type AlibabaAlihouseNewhomeReviewSyncRequest struct {
    model.Params

    // 测评草稿信息
    review   *ProjectReviewDraftDto 

}

func NewAlibabaAlihouseNewhomeReviewSyncRequest() *AlibabaAlihouseNewhomeReviewSyncRequest{
    return &AlibabaAlihouseNewhomeReviewSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihouseNewhomeReviewSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.review.sync"
}

func (r AlibabaAlihouseNewhomeReviewSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihouseNewhomeReviewSyncRequest) SetReview(review *ProjectReviewDraftDto) error {
    r.review = review
    r.Set("review", review)
    return nil
}

func (r AlibabaAlihouseNewhomeReviewSyncRequest) GetReview() *ProjectReviewDraftDto {
    return r.review
}

