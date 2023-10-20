package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeReviewSyncAPIRequest 天猫好房楼盘评测同步 API请求
// alibaba.alihouse.newhome.review.sync
//
// 接受楼盘测评图文信息
type AlibabaAlihouseNewhomeReviewSyncAPIRequest struct {
	model.Params
	// 测评草稿信息
	_review *ProjectReviewDraftDto
}

// NewAlibabaAlihouseNewhomeReviewSyncRequest 初始化AlibabaAlihouseNewhomeReviewSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeReviewSyncRequest() *AlibabaAlihouseNewhomeReviewSyncAPIRequest {
	return &AlibabaAlihouseNewhomeReviewSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeReviewSyncAPIRequest) Reset() {
	r._review = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeReviewSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeReviewSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeReviewSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReview is Review Setter
// 测评草稿信息
func (r *AlibabaAlihouseNewhomeReviewSyncAPIRequest) SetReview(_review *ProjectReviewDraftDto) error {
	r._review = _review
	r.Set("review", _review)
	return nil
}

// GetReview Review Getter
func (r AlibabaAlihouseNewhomeReviewSyncAPIRequest) GetReview() *ProjectReviewDraftDto {
	return r._review
}

var poolAlibabaAlihouseNewhomeReviewSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeReviewSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeReviewSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeReviewSyncAPIRequest
func GetAlibabaAlihouseNewhomeReviewSyncAPIRequest() *AlibabaAlihouseNewhomeReviewSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeReviewSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeReviewSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeReviewSyncAPIRequest 将 AlibabaAlihouseNewhomeReviewSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeReviewSyncAPIRequest(v *AlibabaAlihouseNewhomeReviewSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeReviewSyncAPIRequest.Put(v)
}
