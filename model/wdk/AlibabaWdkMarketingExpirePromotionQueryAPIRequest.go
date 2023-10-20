package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingExpirePromotionQueryAPIRequest) Reset() {
	r._shopId = ""
	r._skuCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingExpirePromotionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingExpirePromotionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingExpirePromotionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkMarketingExpirePromotionQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingExpirePromotionQueryRequest()
	},
}

// GetAlibabaWdkMarketingExpirePromotionQueryRequest 从 sync.Pool 获取 AlibabaWdkMarketingExpirePromotionQueryAPIRequest
func GetAlibabaWdkMarketingExpirePromotionQueryAPIRequest() *AlibabaWdkMarketingExpirePromotionQueryAPIRequest {
	return poolAlibabaWdkMarketingExpirePromotionQueryAPIRequest.Get().(*AlibabaWdkMarketingExpirePromotionQueryAPIRequest)
}

// ReleaseAlibabaWdkMarketingExpirePromotionQueryAPIRequest 将 AlibabaWdkMarketingExpirePromotionQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingExpirePromotionQueryAPIRequest(v *AlibabaWdkMarketingExpirePromotionQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingExpirePromotionQueryAPIRequest.Put(v)
}
