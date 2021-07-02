package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMarketingPromotionKfcAPIRequest 定向优惠活动名称与描述违禁词检查 API请求
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
type TaobaoMarketingPromotionKfcAPIRequest struct {
	model.Params
	// 活动名称
	_promotionTitle string
	// 活动描述
	_promotionDesc string
}

// NewTaobaoMarketingPromotionKfcRequest 初始化TaobaoMarketingPromotionKfcAPIRequest对象
func NewTaobaoMarketingPromotionKfcRequest() *TaobaoMarketingPromotionKfcAPIRequest {
	return &TaobaoMarketingPromotionKfcAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMarketingPromotionKfcAPIRequest) GetApiMethodName() string {
	return "taobao.marketing.promotion.kfc"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMarketingPromotionKfcAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPromotionTitle is PromotionTitle Setter
// 活动名称
func (r *TaobaoMarketingPromotionKfcAPIRequest) SetPromotionTitle(_promotionTitle string) error {
	r._promotionTitle = _promotionTitle
	r.Set("promotion_title", _promotionTitle)
	return nil
}

// GetPromotionTitle PromotionTitle Getter
func (r TaobaoMarketingPromotionKfcAPIRequest) GetPromotionTitle() string {
	return r._promotionTitle
}

// SetPromotionDesc is PromotionDesc Setter
// 活动描述
func (r *TaobaoMarketingPromotionKfcAPIRequest) SetPromotionDesc(_promotionDesc string) error {
	r._promotionDesc = _promotionDesc
	r.Set("promotion_desc", _promotionDesc)
	return nil
}

// GetPromotionDesc PromotionDesc Getter
func (r TaobaoMarketingPromotionKfcAPIRequest) GetPromotionDesc() string {
	return r._promotionDesc
}
