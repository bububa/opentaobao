package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitembuygiftadditemAPIRequest 增加买赠活动商品。【注意，此接口暂不支持并发！】 API请求
// alibaba.wdk.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
type AlibabawdkmarketingitembuygiftadditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemBuyGiftSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingitembuygiftadditemRequest 初始化AlibabawdkmarketingitembuygiftadditemAPIRequest对象
func NewAlibabawdkmarketingitembuygiftadditemRequest() *AlibabawdkmarketingitembuygiftadditemAPIRequest {
	return &AlibabawdkmarketingitembuygiftadditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitembuygiftadditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitembuygiftadditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitembuygiftadditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabawdkmarketingitembuygiftadditemAPIRequest) SetParam0(_param0 *ItemBuyGiftSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitembuygiftadditemAPIRequest) GetParam0() *ItemBuyGiftSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabawdkmarketingitembuygiftadditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingitembuygiftadditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
