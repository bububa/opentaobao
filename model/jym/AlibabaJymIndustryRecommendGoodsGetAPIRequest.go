package jym

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymIndustryRecommendGoodsGetAPIRequest 获取交易猫推荐商品 API请求
// alibaba.jym.industry.recommend.goods.get
//
// 获取交易猫推荐商品
type AlibabaJymIndustryRecommendGoodsGetAPIRequest struct {
	model.Params
	// 请求
	_jymRecommendGoodsRequestDto *JymRecommendGoodsRequestDto
}

// NewAlibabaJymIndustryRecommendGoodsGetRequest 初始化AlibabaJymIndustryRecommendGoodsGetAPIRequest对象
func NewAlibabaJymIndustryRecommendGoodsGetRequest() *AlibabaJymIndustryRecommendGoodsGetAPIRequest {
	return &AlibabaJymIndustryRecommendGoodsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymIndustryRecommendGoodsGetAPIRequest) Reset() {
	r._jymRecommendGoodsRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymIndustryRecommendGoodsGetAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.industry.recommend.goods.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymIndustryRecommendGoodsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymIndustryRecommendGoodsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetJymRecommendGoodsRequestDto is JymRecommendGoodsRequestDto Setter
// 请求
func (r *AlibabaJymIndustryRecommendGoodsGetAPIRequest) SetJymRecommendGoodsRequestDto(_jymRecommendGoodsRequestDto *JymRecommendGoodsRequestDto) error {
	r._jymRecommendGoodsRequestDto = _jymRecommendGoodsRequestDto
	r.Set("jym_recommend_goods_request_dto", _jymRecommendGoodsRequestDto)
	return nil
}

// GetJymRecommendGoodsRequestDto JymRecommendGoodsRequestDto Getter
func (r AlibabaJymIndustryRecommendGoodsGetAPIRequest) GetJymRecommendGoodsRequestDto() *JymRecommendGoodsRequestDto {
	return r._jymRecommendGoodsRequestDto
}

var poolAlibabaJymIndustryRecommendGoodsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymIndustryRecommendGoodsGetRequest()
	},
}

// GetAlibabaJymIndustryRecommendGoodsGetRequest 从 sync.Pool 获取 AlibabaJymIndustryRecommendGoodsGetAPIRequest
func GetAlibabaJymIndustryRecommendGoodsGetAPIRequest() *AlibabaJymIndustryRecommendGoodsGetAPIRequest {
	return poolAlibabaJymIndustryRecommendGoodsGetAPIRequest.Get().(*AlibabaJymIndustryRecommendGoodsGetAPIRequest)
}

// ReleaseAlibabaJymIndustryRecommendGoodsGetAPIRequest 将 AlibabaJymIndustryRecommendGoodsGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymIndustryRecommendGoodsGetAPIRequest(v *AlibabaJymIndustryRecommendGoodsGetAPIRequest) {
	v.Reset()
	poolAlibabaJymIndustryRecommendGoodsGetAPIRequest.Put(v)
}
