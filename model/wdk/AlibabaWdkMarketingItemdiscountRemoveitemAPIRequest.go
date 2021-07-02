package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest 移除报名的商品 API请求
// alibaba.wdk.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
type AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemDiscountSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingItemdiscountRemoveitemRequest 初始化AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountRemoveitemRequest() *AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest {
	return &AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest) SetParam0(_param0 *ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest) GetParam0() *ItemDiscountSku {
	return r._param0
}

// Set is Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// Get Param1 Getter
func (r AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
