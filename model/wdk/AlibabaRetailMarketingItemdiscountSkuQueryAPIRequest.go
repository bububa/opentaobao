package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest 查询单品特价活动商品【同城零售】 API请求
// alibaba.retail.marketing.itemdiscount.sku.query
//
// 查询单品特价活动商品【同城零售】
type AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest struct {
	model.Params
	// 请求体
	_param0 *ItemDiscountActivitySkuQueryRequest
}

// NewAlibabaRetailMarketingItemdiscountSkuQueryRequest 初始化AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountSkuQueryRequest() *AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest {
	return &AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求体
func (r *AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest) SetParam0(_param0 *ItemDiscountActivitySkuQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest) GetParam0() *ItemDiscountActivitySkuQueryRequest {
	return r._param0
}
