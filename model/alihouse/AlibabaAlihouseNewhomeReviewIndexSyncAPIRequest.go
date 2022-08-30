package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest 新测评乐居指数接口 API请求
// alibaba.alihouse.newhome.review.index.sync
//
// 新测评乐居指数同步数据
type AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest struct {
	model.Params
	// 实体类
	_syncProjectNewReviewIndexDto *SyncProjectNewReviewIndexDto
}

// NewAlibabaAlihouseNewhomeReviewIndexSyncRequest 初始化AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeReviewIndexSyncRequest() *AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest {
	return &AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.review.index.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSyncProjectNewReviewIndexDto is SyncProjectNewReviewIndexDto Setter
// 实体类
func (r *AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) SetSyncProjectNewReviewIndexDto(_syncProjectNewReviewIndexDto *SyncProjectNewReviewIndexDto) error {
	r._syncProjectNewReviewIndexDto = _syncProjectNewReviewIndexDto
	r.Set("sync_project_new_review_index_dto", _syncProjectNewReviewIndexDto)
	return nil
}

// GetSyncProjectNewReviewIndexDto SyncProjectNewReviewIndexDto Getter
func (r AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) GetSyncProjectNewReviewIndexDto() *SyncProjectNewReviewIndexDto {
	return r._syncProjectNewReviewIndexDto
}
