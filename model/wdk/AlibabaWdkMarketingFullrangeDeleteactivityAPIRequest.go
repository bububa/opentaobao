package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest 全场活动删除活动接口 API请求
// alibaba.wdk.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
type AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabaWdkMarketingFullrangeDeleteactivityRequest 初始化AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest对象
func NewAlibabaWdkMarketingFullrangeDeleteactivityRequest() *AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest {
	return &AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
