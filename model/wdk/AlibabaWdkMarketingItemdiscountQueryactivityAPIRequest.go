package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest 查找特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingItemdiscountQueryactivityRequest 初始化AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingItemdiscountQueryactivityRequest() *AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest {
	return &AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
