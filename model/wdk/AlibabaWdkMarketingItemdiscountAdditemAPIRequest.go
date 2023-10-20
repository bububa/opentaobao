package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitemdiscountadditemAPIRequest 报名特价商品 API请求
// alibaba.wdk.marketing.itemdiscount.additem
//
// 在商品特价活动中报名特价商品
type AlibabawdkmarketingitemdiscountadditemAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemDiscountSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingitemdiscountadditemRequest 初始化AlibabawdkmarketingitemdiscountadditemAPIRequest对象
func NewAlibabawdkmarketingitemdiscountadditemRequest() *AlibabawdkmarketingitemdiscountadditemAPIRequest {
	return &AlibabawdkmarketingitemdiscountadditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitemdiscountadditemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.additem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitemdiscountadditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitemdiscountadditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabawdkmarketingitemdiscountadditemAPIRequest) SetParam0(_param0 *ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingitemdiscountadditemAPIRequest) GetParam0() *ItemDiscountSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabawdkmarketingitemdiscountadditemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingitemdiscountadditemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
