package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeQueryactivityAPIRequest 全场活动查询活动 API请求
// alibaba.wdk.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabaWdkMarketingFullrangeQueryactivityAPIRequest struct {
	model.Params
	// 查询活动入参
	_param0 *CommonActivityParam
}

// NewAlibabaWdkMarketingFullrangeQueryactivityRequest 初始化AlibabaWdkMarketingFullrangeQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingFullrangeQueryactivityRequest() *AlibabaWdkMarketingFullrangeQueryactivityAPIRequest {
	return &AlibabaWdkMarketingFullrangeQueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 查询活动入参
func (r *AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaWdkMarketingFullrangeQueryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
