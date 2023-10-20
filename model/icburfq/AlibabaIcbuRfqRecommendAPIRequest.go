package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicburfqrecommendAPIRequest rfq推荐 API请求
// alibaba.icbu.rfq.recommend
//
// rfq推荐
type AlibabaicburfqrecommendAPIRequest struct {
	model.Params
	// 入参数据
	_queryDto *QueryDto
}

// NewAlibabaicburfqrecommendRequest 初始化AlibabaicburfqrecommendAPIRequest对象
func NewAlibabaicburfqrecommendRequest() *AlibabaicburfqrecommendAPIRequest {
	return &AlibabaicburfqrecommendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicburfqrecommendAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.rfq.recommend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicburfqrecommendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicburfqrecommendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// 入参数据
func (r *AlibabaicburfqrecommendAPIRequest) SetQueryDto(_queryDto *QueryDto) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaicburfqrecommendAPIRequest) GetQueryDto() *QueryDto {
	return r._queryDto
}
