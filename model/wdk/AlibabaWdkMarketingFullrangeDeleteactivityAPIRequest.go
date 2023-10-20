package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingfullrangedeleteactivityAPIRequest 全场活动删除活动接口 API请求
// alibaba.wdk.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
type AlibabawdkmarketingfullrangedeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabawdkmarketingfullrangedeleteactivityRequest 初始化AlibabawdkmarketingfullrangedeleteactivityAPIRequest对象
func NewAlibabawdkmarketingfullrangedeleteactivityRequest() *AlibabawdkmarketingfullrangedeleteactivityAPIRequest {
	return &AlibabawdkmarketingfullrangedeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingfullrangedeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingfullrangedeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingfullrangedeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabawdkmarketingfullrangedeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingfullrangedeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
