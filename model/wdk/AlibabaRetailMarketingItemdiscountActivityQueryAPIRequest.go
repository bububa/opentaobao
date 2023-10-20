package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest 查询单品特价活动【同城零售】 API请求
// alibaba.retail.marketing.itemdiscount.activity.query
//
// 查询单品特价活动【同城零售】
type AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest struct {
	model.Params
	// 请求体
	_param0 *ItemDiscountActivityQueryRequest
}

// NewAlibabaRetailMarketingItemdiscountActivityQueryRequest 初始化AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest对象
func NewAlibabaRetailMarketingItemdiscountActivityQueryRequest() *AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest {
	return &AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itemdiscount.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求体
func (r *AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest) SetParam0(_param0 *ItemDiscountActivityQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest) GetParam0() *ItemDiscountActivityQueryRequest {
	return r._param0
}
