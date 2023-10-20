package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest 全场增加换购品 API请求
// alibaba.hm.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemStairSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingFullrangeAddexchangeitemRequest 初始化AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest对象
func NewAlibabaHmMarketingFullrangeAddexchangeitemRequest() *AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest {
	return &AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.addexchangeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest) GetParam0() *ItemStairSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingFullrangeAddexchangeitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
