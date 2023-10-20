package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordlistrelevantproductsAPIRequest 查询和词匹配的推广产品 API请求
// alibaba.scbp.ad.keyword.list.relevant.products
//
// 查询和词匹配的推广产品
type AlibabascbpadkeywordlistrelevantproductsAPIRequest struct {
	model.Params
	// 关键词
	_keyword string
	// 用户信息
	_topContext *TopContextDto
	// 计划id
	_campaignId int64
}

// NewAlibabascbpadkeywordlistrelevantproductsRequest 初始化AlibabascbpadkeywordlistrelevantproductsAPIRequest对象
func NewAlibabascbpadkeywordlistrelevantproductsRequest() *AlibabascbpadkeywordlistrelevantproductsAPIRequest {
	return &AlibabascbpadkeywordlistrelevantproductsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordlistrelevantproductsAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.list.relevant.products"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordlistrelevantproductsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordlistrelevantproductsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键词
func (r *AlibabascbpadkeywordlistrelevantproductsAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlibabascbpadkeywordlistrelevantproductsAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetTopContext is TopContext Setter
// 用户信息
func (r *AlibabascbpadkeywordlistrelevantproductsAPIRequest) SetTopContext(_topContext *TopContextDto) error {
	r._topContext = _topContext
	r.Set("top_context", _topContext)
	return nil
}

// GetTopContext TopContext Getter
func (r AlibabascbpadkeywordlistrelevantproductsAPIRequest) GetTopContext() *TopContextDto {
	return r._topContext
}

// SetCampaignId is CampaignId Setter
// 计划id
func (r *AlibabascbpadkeywordlistrelevantproductsAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r AlibabascbpadkeywordlistrelevantproductsAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
