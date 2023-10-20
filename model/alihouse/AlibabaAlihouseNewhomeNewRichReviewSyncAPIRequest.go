package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest 新评测改造接口同步 API请求
// alibaba.alihouse.newhome.new.rich.review.sync
//
// 新评测改造接口同步
type AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest struct {
	model.Params
	// 评测实体类
	_syncRichReviewDto *SyncRichReviewDto
}

// NewAlibabaAlihouseNewhomeNewRichReviewSyncRequest 初始化AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeNewRichReviewSyncRequest() *AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest {
	return &AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.new.rich.review.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncRichReviewDto is SyncRichReviewDto Setter
// 评测实体类
func (r *AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest) SetSyncRichReviewDto(_syncRichReviewDto *SyncRichReviewDto) error {
	r._syncRichReviewDto = _syncRichReviewDto
	r.Set("sync_rich_review_dto", _syncRichReviewDto)
	return nil
}

// GetSyncRichReviewDto SyncRichReviewDto Getter
func (r AlibabaAlihouseNewhomeNewRichReviewSyncAPIRequest) GetSyncRichReviewDto() *SyncRichReviewDto {
	return r._syncRichReviewDto
}
