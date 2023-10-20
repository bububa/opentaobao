package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomenewreviewsyncAPIRequest 新测评基础信息接口 API请求
// alibaba.alihouse.newhome.new.review.sync
//
// 新测评基础信息接口
type AlibabaalihousenewhomenewreviewsyncAPIRequest struct {
	model.Params
	// 测评数据实体类
	_syncProjectNewReviewDraftDto *SyncProjectNewReviewDraftDto
}

// NewAlibabaalihousenewhomenewreviewsyncRequest 初始化AlibabaalihousenewhomenewreviewsyncAPIRequest对象
func NewAlibabaalihousenewhomenewreviewsyncRequest() *AlibabaalihousenewhomenewreviewsyncAPIRequest {
	return &AlibabaalihousenewhomenewreviewsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomenewreviewsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.new.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomenewreviewsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomenewreviewsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncProjectNewReviewDraftDto is SyncProjectNewReviewDraftDto Setter
// 测评数据实体类
func (r *AlibabaalihousenewhomenewreviewsyncAPIRequest) SetSyncProjectNewReviewDraftDto(_syncProjectNewReviewDraftDto *SyncProjectNewReviewDraftDto) error {
	r._syncProjectNewReviewDraftDto = _syncProjectNewReviewDraftDto
	r.Set("sync_project_new_review_draft_dto", _syncProjectNewReviewDraftDto)
	return nil
}

// GetSyncProjectNewReviewDraftDto SyncProjectNewReviewDraftDto Getter
func (r AlibabaalihousenewhomenewreviewsyncAPIRequest) GetSyncProjectNewReviewDraftDto() *SyncProjectNewReviewDraftDto {
	return r._syncProjectNewReviewDraftDto
}
