package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingFullrangeCreateactivityAPIRequest 创建全场活动 API请求
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabaWdkMarketingFullrangeCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *FullRangeActivity
}

// NewAlibabaWdkMarketingFullrangeCreateactivityRequest 初始化AlibabaWdkMarketingFullrangeCreateactivityAPIRequest对象
func NewAlibabaWdkMarketingFullrangeCreateactivityRequest() *AlibabaWdkMarketingFullrangeCreateactivityAPIRequest {
	return &AlibabaWdkMarketingFullrangeCreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) SetParam(_param *FullRangeActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkMarketingFullrangeCreateactivityAPIRequest) GetParam() *FullRangeActivity {
	return r._param
}
