package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest 全场增加换购品 API请求
// alibaba.wdk.marketing.fullrange.addexchangeitem
//
// 全场增加换购品
type AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *ItemStairSku
	// 系统自动生成
	_param1 *CommonActivityParam
}

// NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest 初始化AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest对象
func NewAlibabaWdkMarketingFullrangeAddexchangeitemRequest() *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest {
	return &AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.addexchangeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) SetParam0(_param0 *ItemStairSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetParam0() *ItemStairSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaWdkMarketingFullrangeAddexchangeitemAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}
