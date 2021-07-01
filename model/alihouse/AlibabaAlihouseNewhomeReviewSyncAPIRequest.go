package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeReviewSyncAPIRequest
天猫好房楼盘评测同步 API请求
alibaba.alihouse.newhome.review.sync

接受楼盘测评图文信息 */
type AlibabaAlihouseNewhomeReviewSyncAPIRequest struct {
	model.Params
	// 测评草稿信息
	_review *ProjectReviewDraftDto
}

// NewAlibabaAlihouseNewhomeReviewSyncRequest 初始化AlibabaAlihouseNewhomeReviewSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeReviewSyncRequest() *AlibabaAlihouseNewhomeReviewSyncAPIRequest {
	return &AlibabaAlihouseNewhomeReviewSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeReviewSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeReviewSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Review Setter
// 测评草稿信息
func (r *AlibabaAlihouseNewhomeReviewSyncAPIRequest) SetReview(_review *ProjectReviewDraftDto) error {
	r._review = _review
	r.Set("review", _review)
	return nil
}

// Get Review Getter
func (r AlibabaAlihouseNewhomeReviewSyncAPIRequest) GetReview() *ProjectReviewDraftDto {
	return r._review
}
