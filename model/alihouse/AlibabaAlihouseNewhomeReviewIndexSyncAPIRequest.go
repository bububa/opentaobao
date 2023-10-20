package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) Reset() {
	r._syncProjectNewReviewIndexDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.review.index.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeReviewIndexSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeReviewIndexSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeReviewIndexSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest
func GetAlibabaAlihouseNewhomeReviewIndexSyncAPIRequest() *AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeReviewIndexSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeReviewIndexSyncAPIRequest 将 AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeReviewIndexSyncAPIRequest(v *AlibabaAlihouseNewhomeReviewIndexSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeReviewIndexSyncAPIRequest.Put(v)
}
