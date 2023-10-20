package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpPromotionGlobalDiscountGetAPIRequest 获取卖家最低折扣 API请求
// taobao.ump.promotion.global.discount.get
//
// 提供卖家最低折扣查询功能
type TaobaoUmpPromotionGlobalDiscountGetAPIRequest struct {
	model.Params
}

// NewTaobaoUmpPromotionGlobalDiscountGetRequest 初始化TaobaoUmpPromotionGlobalDiscountGetAPIRequest对象
func NewTaobaoUmpPromotionGlobalDiscountGetRequest() *TaobaoUmpPromotionGlobalDiscountGetAPIRequest {
	return &TaobaoUmpPromotionGlobalDiscountGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpPromotionGlobalDiscountGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.promotion.global.discount.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpPromotionGlobalDiscountGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpPromotionGlobalDiscountGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
