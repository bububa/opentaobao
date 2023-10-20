package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapricepromotionitemaddAPIRequest 新增档期商品 API请求
// alibaba.price.promotion.item.add
//
// 批量新增档期活动商品
type AlibabapricepromotionitemaddAPIRequest struct {
	model.Params
	// 入参
	_promotionContent *PromotionContent
}

// NewAlibabapricepromotionitemaddRequest 初始化AlibabapricepromotionitemaddAPIRequest对象
func NewAlibabapricepromotionitemaddRequest() *AlibabapricepromotionitemaddAPIRequest {
	return &AlibabapricepromotionitemaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapricepromotionitemaddAPIRequest) GetApiMethodName() string {
	return "alibaba.price.promotion.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapricepromotionitemaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapricepromotionitemaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPromotionContent is PromotionContent Setter
// 入参
func (r *AlibabapricepromotionitemaddAPIRequest) SetPromotionContent(_promotionContent *PromotionContent) error {
	r._promotionContent = _promotionContent
	r.Set("promotion_content", _promotionContent)
	return nil
}

// GetPromotionContent PromotionContent Getter
func (r AlibabapricepromotionitemaddAPIRequest) GetPromotionContent() *PromotionContent {
	return r._promotionContent
}
