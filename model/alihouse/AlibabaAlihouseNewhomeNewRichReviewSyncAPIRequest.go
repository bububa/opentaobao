package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomenewrichreviewsyncAPIRequest 新评测改造接口同步 API请求
// alibaba.alihouse.newhome.new.rich.review.sync
//
// 新评测改造接口同步
type AlibabaalihousenewhomenewrichreviewsyncAPIRequest struct {
	model.Params
	// 评测实体类
	_syncRichReviewDto *SyncRichReviewDto
}

// NewAlibabaalihousenewhomenewrichreviewsyncRequest 初始化AlibabaalihousenewhomenewrichreviewsyncAPIRequest对象
func NewAlibabaalihousenewhomenewrichreviewsyncRequest() *AlibabaalihousenewhomenewrichreviewsyncAPIRequest {
	return &AlibabaalihousenewhomenewrichreviewsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomenewrichreviewsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.new.rich.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomenewrichreviewsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomenewrichreviewsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncRichReviewDto is SyncRichReviewDto Setter
// 评测实体类
func (r *AlibabaalihousenewhomenewrichreviewsyncAPIRequest) SetSyncRichReviewDto(_syncRichReviewDto *SyncRichReviewDto) error {
	r._syncRichReviewDto = _syncRichReviewDto
	r.Set("sync_rich_review_dto", _syncRichReviewDto)
	return nil
}

// GetSyncRichReviewDto SyncRichReviewDto Getter
func (r AlibabaalihousenewhomenewrichreviewsyncAPIRequest) GetSyncRichReviewDto() *SyncRichReviewDto {
	return r._syncRichReviewDto
}
