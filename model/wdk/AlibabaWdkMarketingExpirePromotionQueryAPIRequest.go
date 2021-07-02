package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingExpirePromotionQueryAPIRequest 短保优惠查询 API请求
// alibaba.wdk.marketing.expire.promotion.query
//
// 短保优惠查询
type AlibabaWdkMarketingExpirePromotionQueryAPIRequest struct {
	model.Params
	// 店铺id
	_shopId string
	// 商品skucode
	_skuCode string
}

// NewAlibabaWdkMarketingExpirePromotionQueryRequest 初始化AlibabaWdkMarketingExpirePromotionQueryAPIRequest对象
func NewAlibabaWdkMarketingExpirePromotionQueryRequest() *AlibabaWdkMarketingExpirePromotionQueryAPIRequest {
	return &AlibabaWdkMarketingExpirePromotionQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetShopId is ShopId Setter
// 店铺id
func (r *AlibabaWdkMarketingExpirePromotionQueryAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaWdkMarketingExpirePromotionQueryAPIRequest) GetShopId() string {
	return r._shopId
}

// SetSkuCode is SkuCode Setter
// 商品skucode
func (r *AlibabaWdkMarketingExpirePromotionQueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabaWdkMarketingExpirePromotionQueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}
