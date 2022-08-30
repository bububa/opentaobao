package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.new.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeNewReviewSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
