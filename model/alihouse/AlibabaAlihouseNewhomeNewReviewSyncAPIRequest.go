package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeNewReviewSyncAPIRequest 新测评基础信息接口 API请求
// alibaba.alihouse.newhome.new.review.sync
//
// 新测评基础信息接口
type AlibabaAlihouseNewhomeNewReviewSyncAPIRequest struct {
	model.Params
	// 测评数据实体类
	_syncProjectNewReviewDraftDto *SyncProjectNewReviewDraftDto
}

// NewAlibabaAlihouseNewhomeNewReviewSyncRequest 初始化AlibabaAlihouseNewhomeNewReviewSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeNewReviewSyncRequest() *AlibabaAlihouseNewhomeNewReviewSyncAPIRequest {
	return &AlibabaAlihouseNewhomeNewReviewSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) Reset() {
	r._syncProjectNewReviewDraftDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.new.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncProjectNewReviewDraftDto is SyncProjectNewReviewDraftDto Setter
// 测评数据实体类
func (r *AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) SetSyncProjectNewReviewDraftDto(_syncProjectNewReviewDraftDto *SyncProjectNewReviewDraftDto) error {
	r._syncProjectNewReviewDraftDto = _syncProjectNewReviewDraftDto
	r.Set("sync_project_new_review_draft_dto", _syncProjectNewReviewDraftDto)
	return nil
}

// GetSyncProjectNewReviewDraftDto SyncProjectNewReviewDraftDto Getter
func (r AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) GetSyncProjectNewReviewDraftDto() *SyncProjectNewReviewDraftDto {
	return r._syncProjectNewReviewDraftDto
}

var poolAlibabaAlihouseNewhomeNewReviewSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeNewReviewSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeNewReviewSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeNewReviewSyncAPIRequest
func GetAlibabaAlihouseNewhomeNewReviewSyncAPIRequest() *AlibabaAlihouseNewhomeNewReviewSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeNewReviewSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeNewReviewSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeNewReviewSyncAPIRequest 将 AlibabaAlihouseNewhomeNewReviewSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeNewReviewSyncAPIRequest(v *AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeNewReviewSyncAPIRequest.Put(v)
}
