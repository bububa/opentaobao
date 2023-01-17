package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest 查询买赠活动 API请求
// alibaba.wdk.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest struct {
	model.Params
	// 查询入参
	_param *CommonActivityParam
}

// NewAlibabaWdkMarketingItembuygiftQueryactivityRequest 初始化AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest对象
func NewAlibabaWdkMarketingItembuygiftQueryactivityRequest() *AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest {
	return &AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
