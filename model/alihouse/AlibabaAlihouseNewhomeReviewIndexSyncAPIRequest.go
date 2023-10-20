package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomereviewindexsyncAPIRequest 新测评乐居指数接口 API请求
// alibaba.alihouse.newhome.review.index.sync
//
// 新测评乐居指数同步数据
type AlibabaalihousenewhomereviewindexsyncAPIRequest struct {
	model.Params
	// 实体类
	_syncProjectNewReviewIndexDto *SyncProjectNewReviewIndexDto
}

// NewAlibabaalihousenewhomereviewindexsyncRequest 初始化AlibabaalihousenewhomereviewindexsyncAPIRequest对象
func NewAlibabaalihousenewhomereviewindexsyncRequest() *AlibabaalihousenewhomereviewindexsyncAPIRequest {
	return &AlibabaalihousenewhomereviewindexsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomereviewindexsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.review.index.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomereviewindexsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomereviewindexsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncProjectNewReviewIndexDto is SyncProjectNewReviewIndexDto Setter
// 实体类
func (r *AlibabaalihousenewhomereviewindexsyncAPIRequest) SetSyncProjectNewReviewIndexDto(_syncProjectNewReviewIndexDto *SyncProjectNewReviewIndexDto) error {
	r._syncProjectNewReviewIndexDto = _syncProjectNewReviewIndexDto
	r.Set("sync_project_new_review_index_dto", _syncProjectNewReviewIndexDto)
	return nil
}

// GetSyncProjectNewReviewIndexDto SyncProjectNewReviewIndexDto Getter
func (r AlibabaalihousenewhomereviewindexsyncAPIRequest) GetSyncProjectNewReviewIndexDto() *SyncProjectNewReviewIndexDto {
	return r._syncProjectNewReviewIndexDto
}
