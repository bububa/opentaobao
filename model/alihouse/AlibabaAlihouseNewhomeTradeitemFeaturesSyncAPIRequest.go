package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometradeitemfeaturessyncAPIRequest 同步品活动标 API请求
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
type AlibabaalihousenewhometradeitemfeaturessyncAPIRequest struct {
	model.Params
	// 请求对象
	_tagsRequestDto *TagsRequestDto
}

// NewAlibabaalihousenewhometradeitemfeaturessyncRequest 初始化AlibabaalihousenewhometradeitemfeaturessyncAPIRequest对象
func NewAlibabaalihousenewhometradeitemfeaturessyncRequest() *AlibabaalihousenewhometradeitemfeaturessyncAPIRequest {
	return &AlibabaalihousenewhometradeitemfeaturessyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhometradeitemfeaturessyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.tradeitem.features.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhometradeitemfeaturessyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhometradeitemfeaturessyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagsRequestDto is TagsRequestDto Setter
// 请求对象
func (r *AlibabaalihousenewhometradeitemfeaturessyncAPIRequest) SetTagsRequestDto(_tagsRequestDto *TagsRequestDto) error {
	r._tagsRequestDto = _tagsRequestDto
	r.Set("tags_request_dto", _tagsRequestDto)
	return nil
}

// GetTagsRequestDto TagsRequestDto Getter
func (r AlibabaalihousenewhometradeitemfeaturessyncAPIRequest) GetTagsRequestDto() *TagsRequestDto {
	return r._tagsRequestDto
}
