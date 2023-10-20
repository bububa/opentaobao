package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomarketingpromotionkfcAPIRequest 定向优惠活动名称与描述违禁词检查 API请求
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
type TaobaomarketingpromotionkfcAPIRequest struct {
	model.Params
	// 活动名称
	_promotionTitle string
	// 活动描述
	_promotionDesc string
}

// NewTaobaomarketingpromotionkfcRequest 初始化TaobaomarketingpromotionkfcAPIRequest对象
func NewTaobaomarketingpromotionkfcRequest() *TaobaomarketingpromotionkfcAPIRequest {
	return &TaobaomarketingpromotionkfcAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomarketingpromotionkfcAPIRequest) GetApiMethodName() string {
	return "taobao.marketing.promotion.kfc"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomarketingpromotionkfcAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomarketingpromotionkfcAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionTitle is PromotionTitle Setter
// 活动名称
func (r *TaobaomarketingpromotionkfcAPIRequest) SetPromotionTitle(_promotionTitle string) error {
	r._promotionTitle = _promotionTitle
	r.Set("promotion_title", _promotionTitle)
	return nil
}

// GetPromotionTitle PromotionTitle Getter
func (r TaobaomarketingpromotionkfcAPIRequest) GetPromotionTitle() string {
	return r._promotionTitle
}

// SetPromotionDesc is PromotionDesc Setter
// 活动描述
func (r *TaobaomarketingpromotionkfcAPIRequest) SetPromotionDesc(_promotionDesc string) error {
	r._promotionDesc = _promotionDesc
	r.Set("promotion_desc", _promotionDesc)
	return nil
}

// GetPromotionDesc PromotionDesc Getter
func (r TaobaomarketingpromotionkfcAPIRequest) GetPromotionDesc() string {
	return r._promotionDesc
}
