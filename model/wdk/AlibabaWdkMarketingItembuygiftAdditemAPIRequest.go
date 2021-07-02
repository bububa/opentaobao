package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftAdditemAPIRequest 增加买赠活动商品。【注意，此接口暂不支持并发！】 API请求
// alibaba.wdk.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
type AlibabaWdkMarketingItembuygiftAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemBuyGiftSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingItembuygiftAdditemRequest 初始化AlibabaWdkMarketingItembuygiftAdditemAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftAdditemRequest() *AlibabaWdkMarketingItembuygiftAdditemAPIRequest {
	return &AlibabaWdkMarketingItembuygiftAdditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftAdditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftAdditemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItembuygiftAdditemAPIRequest) SetParam0(_param0 *ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItembuygiftAdditemAPIRequest) GetParam0() *ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaWdkMarketingItembuygiftAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingItembuygiftAdditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
