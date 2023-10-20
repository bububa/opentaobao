package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingfullrangequeryactivityAPIRequest 全场活动查询活动 API请求
// alibaba.wdk.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabawdkmarketingfullrangequeryactivityAPIRequest struct {
	model.Params
	// 查询活动入参
	_param0 *CommonActivityParam
}

// NewAlibabawdkmarketingfullrangequeryactivityRequest 初始化AlibabawdkmarketingfullrangequeryactivityAPIRequest对象
func NewAlibabawdkmarketingfullrangequeryactivityRequest() *AlibabawdkmarketingfullrangequeryactivityAPIRequest {
	return &AlibabawdkmarketingfullrangequeryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingfullrangequeryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingfullrangequeryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingfullrangequeryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询活动入参
func (r *AlibabawdkmarketingfullrangequeryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingfullrangequeryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
