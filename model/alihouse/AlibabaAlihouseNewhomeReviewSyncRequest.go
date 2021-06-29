package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫好房楼盘评测同步 API请求
alibaba.alihouse.newhome.review.sync

接受楼盘测评图文信息
*/
type AlibabaAlihouseNewhomeReviewSyncRequest struct {
    model.Params
    // 测评草稿信息
    _review   *ProjectReviewDraftDTO
}

// 初始化AlibabaAlihouseNewhomeReviewSyncRequest对象
func NewAlibabaAlihouseNewhomeReviewSyncRequest() *AlibabaAlihouseNewhomeReviewSyncRequest{
    return &AlibabaAlihouseNewhomeReviewSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeReviewSyncRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.review.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeReviewSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Review Setter
// 测评草稿信息
func (r *AlibabaAlihouseNewhomeReviewSyncRequest) SetReview(_review *ProjectReviewDraftDTO) error {
    r._review = _review
    r.Set("review", _review)
    return nil
}

// Review Getter
func (r AlibabaAlihouseNewhomeReviewSyncRequest) GetReview() *ProjectReviewDraftDTO {
    return r._review
}
