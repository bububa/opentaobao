package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItemdiscountQueryactivityAPIRequest 查找特价活动 API请求
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
type AlibabaHmMarketingItemdiscountQueryactivityAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *CommonActivityParam
}

// NewAlibabaHmMarketingItemdiscountQueryactivityRequest 初始化AlibabaHmMarketingItemdiscountQueryactivityAPIRequest对象
func NewAlibabaHmMarketingItemdiscountQueryactivityRequest() *AlibabaHmMarketingItemdiscountQueryactivityAPIRequest {
	return &AlibabaHmMarketingItemdiscountQueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItemdiscountQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itemdiscount.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItemdiscountQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItemdiscountQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItemdiscountQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItemdiscountQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
