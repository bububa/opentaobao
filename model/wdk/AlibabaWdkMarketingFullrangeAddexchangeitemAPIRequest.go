package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest 全场增加换购品 API请求
// alibaba.wdk.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemStairSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabawdkmarketingfullrangeaddexchangeitemRequest 初始化AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest对象
func NewAlibabawdkmarketingfullrangeaddexchangeitemRequest() *AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest {
	return &AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.addexchangeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest) GetParam0() *ItemStairSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabawdkmarketingfullrangeaddexchangeitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
