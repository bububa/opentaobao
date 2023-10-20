package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomereviewsyncAPIRequest 天猫好房楼盘评测同步 API请求
// alibaba.alihouse.newhome.review.sync
//
// 接受楼盘测评图文信息
type AlibabaalihousenewhomereviewsyncAPIRequest struct {
	model.Params
	// 测评草稿信息
	_review *ProjectReviewDraftDto
}

// NewAlibabaalihousenewhomereviewsyncRequest 初始化AlibabaalihousenewhomereviewsyncAPIRequest对象
func NewAlibabaalihousenewhomereviewsyncRequest() *AlibabaalihousenewhomereviewsyncAPIRequest {
	return &AlibabaalihousenewhomereviewsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomereviewsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomereviewsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomereviewsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReview is Review Setter
// 测评草稿信息
func (r *AlibabaalihousenewhomereviewsyncAPIRequest) SetReview(_review *ProjectReviewDraftDto) error {
	r._review = _review
	r.Set("review", _review)
	return nil
}

// GetReview Review Getter
func (r AlibabaalihousenewhomereviewsyncAPIRequest) GetReview() *ProjectReviewDraftDto {
	return r._review
}
