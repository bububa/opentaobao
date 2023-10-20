package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingexpirepromotionqueryAPIRequest 短保优惠查询 API请求
// alibaba.wdk.marketing.expire.promotion.query
//
// 短保优惠查询
type AlibabawdkmarketingexpirepromotionqueryAPIRequest struct {
	model.Params
	// 店铺id
	_shopId string
	// 商品skucode
	_skuCode string
}

// NewAlibabawdkmarketingexpirepromotionqueryRequest 初始化AlibabawdkmarketingexpirepromotionqueryAPIRequest对象
func NewAlibabawdkmarketingexpirepromotionqueryRequest() *AlibabawdkmarketingexpirepromotionqueryAPIRequest {
	return &AlibabawdkmarketingexpirepromotionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingexpirepromotionqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingexpirepromotionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingexpirepromotionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShopId is ShopId Setter
// 店铺id
func (r *AlibabawdkmarketingexpirepromotionqueryAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabawdkmarketingexpirepromotionqueryAPIRequest) GetShopId() string {
	return r._shopId
}

// SetSkuCode is SkuCode Setter
// 商品skucode
func (r *AlibabawdkmarketingexpirepromotionqueryAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkmarketingexpirepromotionqueryAPIRequest) GetSkuCode() string {
	return r._skuCode
}
