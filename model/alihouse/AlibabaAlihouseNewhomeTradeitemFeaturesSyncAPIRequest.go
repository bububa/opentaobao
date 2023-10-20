package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest 同步品活动标 API请求
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
type AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest struct {
	model.Params
	// 请求对象
	_tagsRequestDto *TagsRequestDto
}

// NewAlibabaAlihouseNewhomeTradeitemFeaturesSyncRequest 初始化AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeTradeitemFeaturesSyncRequest() *AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest {
	return &AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.tradeitem.features.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagsRequestDto is TagsRequestDto Setter
// 请求对象
func (r *AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest) SetTagsRequestDto(_tagsRequestDto *TagsRequestDto) error {
	r._tagsRequestDto = _tagsRequestDto
	r.Set("tags_request_dto", _tagsRequestDto)
	return nil
}

// GetTagsRequestDto TagsRequestDto Getter
func (r AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIRequest) GetTagsRequestDto() *TagsRequestDto {
	return r._tagsRequestDto
}
