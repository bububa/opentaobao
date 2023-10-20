package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymindustryrecommendgoodsgetAPIRequest 获取交易猫推荐商品 API请求
// alibaba.jym.industry.recommend.goods.get
//
// 获取交易猫推荐商品
type AlibabajymindustryrecommendgoodsgetAPIRequest struct {
	model.Params
	// 请求
	_jymRecommendGoodsRequestDto *JymRecommendGoodsRequestDto
}

// NewAlibabajymindustryrecommendgoodsgetRequest 初始化AlibabajymindustryrecommendgoodsgetAPIRequest对象
func NewAlibabajymindustryrecommendgoodsgetRequest() *AlibabajymindustryrecommendgoodsgetAPIRequest {
	return &AlibabajymindustryrecommendgoodsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymindustryrecommendgoodsgetAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.recommend.goods.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymindustryrecommendgoodsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymindustryrecommendgoodsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJymRecommendGoodsRequestDto is JymRecommendGoodsRequestDto Setter
// 请求
func (r *AlibabajymindustryrecommendgoodsgetAPIRequest) SetJymRecommendGoodsRequestDto(_jymRecommendGoodsRequestDto *JymRecommendGoodsRequestDto) error {
	r._jymRecommendGoodsRequestDto = _jymRecommendGoodsRequestDto
	r.Set("jym_recommend_goods_request_dto", _jymRecommendGoodsRequestDto)
	return nil
}

// GetJymRecommendGoodsRequestDto JymRecommendGoodsRequestDto Getter
func (r AlibabajymindustryrecommendgoodsgetAPIRequest) GetJymRecommendGoodsRequestDto() *JymRecommendGoodsRequestDto {
	return r._jymRecommendGoodsRequestDto
}
