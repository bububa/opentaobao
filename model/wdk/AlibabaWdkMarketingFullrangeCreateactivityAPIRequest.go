package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingfullrangecreateactivityAPIRequest 创建全场活动 API请求
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabawdkmarketingfullrangecreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *FullRangeActivity
}

// NewAlibabawdkmarketingfullrangecreateactivityRequest 初始化AlibabawdkmarketingfullrangecreateactivityAPIRequest对象
func NewAlibabawdkmarketingfullrangecreateactivityRequest() *AlibabawdkmarketingfullrangecreateactivityAPIRequest {
	return &AlibabawdkmarketingfullrangecreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingfullrangecreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.fullrange.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingfullrangecreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingfullrangecreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabawdkmarketingfullrangecreateactivityAPIRequest) SetParam(_param *FullRangeActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingfullrangecreateactivityAPIRequest) GetParam() *FullRangeActivity {
	return r._param
}
