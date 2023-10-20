package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItembuygiftRemoveitemAPIRequest 移除买赠活动下的商品。【注意，此接口暂不支持并发！】 API请求
// alibaba.hm.marketing.itembuygift.removeitem
//
// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
type AlibabaHmMarketingItembuygiftRemoveitemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemBuyGiftSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItembuygiftRemoveitemRequest 初始化AlibabaHmMarketingItembuygiftRemoveitemAPIRequest对象
func NewAlibabaHmMarketingItembuygiftRemoveitemRequest() *AlibabaHmMarketingItembuygiftRemoveitemAPIRequest {
	return &AlibabaHmMarketingItembuygiftRemoveitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItembuygiftRemoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItembuygiftRemoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItembuygiftRemoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItembuygiftRemoveitemAPIRequest) SetParam0(_param0 *ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItembuygiftRemoveitemAPIRequest) GetParam0() *ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaHmMarketingItembuygiftRemoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItembuygiftRemoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
