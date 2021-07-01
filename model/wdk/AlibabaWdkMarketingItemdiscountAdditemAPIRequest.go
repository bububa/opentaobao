package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingItemdiscountAdditemAPIRequest
报名特价商品 API请求
alibaba.wdk.marketing.itemdiscount.additem

在商品特价活动中报名特价商品 */
type AlibabaWdkMarketingItemdiscountAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemDiscountSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingItemdiscountAdditemRequest 初始化AlibabaWdkMarketingItemdiscountAdditemAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountAdditemRequest() *AlibabaWdkMarketingItemdiscountAdditemAPIRequest {
	return &AlibabaWdkMarketingItemdiscountAdditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountAdditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountAdditemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItemdiscountAdditemAPIRequest) SetParam0(_param0 *ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingItemdiscountAdditemAPIRequest) GetParam0() *ItemDiscountSku {
	return r._param0
}

// Set is Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItemdiscountAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaWdkMarketingItemdiscountAdditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
