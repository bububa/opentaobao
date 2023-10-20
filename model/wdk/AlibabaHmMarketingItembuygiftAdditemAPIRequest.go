package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftAdditemAPIRequest 增加买赠活动商品。【注意，此接口暂不支持并发！】 API请求
// alibaba.hm.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
type AlibabaHmMarketingItembuygiftAdditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemBuyGiftSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItembuygiftAdditemRequest 初始化AlibabaHmMarketingItembuygiftAdditemAPIRequest对象
func NewAlibabaHmMarketingItembuygiftAdditemRequest() *AlibabaHmMarketingItembuygiftAdditemAPIRequest {
	return &AlibabaHmMarketingItembuygiftAdditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftAdditemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftAdditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftAdditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItembuygiftAdditemAPIRequest) SetParam0(_param0 *ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItembuygiftAdditemAPIRequest) GetParam0() *ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaHmMarketingItembuygiftAdditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItembuygiftAdditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
