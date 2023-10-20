package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingfullrangeaddexchangeitemAPIRequest 全场增加换购品 API请求
// alibaba.hm.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabahmmarketingfullrangeaddexchangeitemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemStairSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabahmmarketingfullrangeaddexchangeitemRequest 初始化AlibabahmmarketingfullrangeaddexchangeitemAPIRequest对象
func NewAlibabahmmarketingfullrangeaddexchangeitemRequest() *AlibabahmmarketingfullrangeaddexchangeitemAPIRequest {
	return &AlibabahmmarketingfullrangeaddexchangeitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingfullrangeaddexchangeitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.addexchangeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingfullrangeaddexchangeitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingfullrangeaddexchangeitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabahmmarketingfullrangeaddexchangeitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingfullrangeaddexchangeitemAPIRequest) GetParam0() *ItemStairSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabahmmarketingfullrangeaddexchangeitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabahmmarketingfullrangeaddexchangeitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
