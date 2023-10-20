package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapricepromotioncreateAPIRequest 营销档期活动创建 API请求
// alibaba.price.promotion.create
//
// 大润发-盒马帮提供新增创建营销活动
type AlibabapricepromotioncreateAPIRequest struct {
	model.Params
	// 档期活动参数
	_promotionActivityDo *PromotionActivityDo
}

// NewAlibabapricepromotioncreateRequest 初始化AlibabapricepromotioncreateAPIRequest对象
func NewAlibabapricepromotioncreateRequest() *AlibabapricepromotioncreateAPIRequest {
	return &AlibabapricepromotioncreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapricepromotioncreateAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapricepromotioncreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapricepromotioncreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionActivityDo is PromotionActivityDo Setter
// 档期活动参数
func (r *AlibabapricepromotioncreateAPIRequest) SetPromotionActivityDo(_promotionActivityDo *PromotionActivityDo) error {
	r._promotionActivityDo = _promotionActivityDo
	r.Set("promotion_activity_do", _promotionActivityDo)
	return nil
}

// GetPromotionActivityDo PromotionActivityDo Getter
func (r AlibabapricepromotioncreateAPIRequest) GetPromotionActivityDo() *PromotionActivityDo {
	return r._promotionActivityDo
}
