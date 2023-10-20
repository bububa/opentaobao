package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingfullrangeremoveitemAPIRequest 全场活动删除购品 API请求
// alibaba.wdk.marketing.fullrange.removeitem
//
// 删除换购商品
type AlibabawdkmarketingfullrangeremoveitemAPIRequest struct {
	model.Params
	// 商品sku信息
	_param0 *ItemStairSku
	// 活动信息
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingfullrangeremoveitemRequest 初始化AlibabawdkmarketingfullrangeremoveitemAPIRequest对象
func NewAlibabawdkmarketingfullrangeremoveitemRequest() *AlibabawdkmarketingfullrangeremoveitemAPIRequest {
	return &AlibabawdkmarketingfullrangeremoveitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingfullrangeremoveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.removeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingfullrangeremoveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingfullrangeremoveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品sku信息
func (r *AlibabawdkmarketingfullrangeremoveitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingfullrangeremoveitemAPIRequest) GetParam0() *ItemStairSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动信息
func (r *AlibabawdkmarketingfullrangeremoveitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingfullrangeremoveitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
