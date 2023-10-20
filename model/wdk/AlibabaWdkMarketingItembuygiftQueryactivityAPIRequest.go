package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitembuygiftqueryactivityAPIRequest 查询买赠活动 API请求
// alibaba.wdk.marketing.itembuygift.queryactivity
//
// 查询买赠活动
type AlibabawdkmarketingitembuygiftqueryactivityAPIRequest struct {
	model.Params
	// 查询入参
	_param *CommonActivityParam
}

// NewAlibabawdkmarketingitembuygiftqueryactivityRequest 初始化AlibabawdkmarketingitembuygiftqueryactivityAPIRequest对象
func NewAlibabawdkmarketingitembuygiftqueryactivityRequest() *AlibabawdkmarketingitembuygiftqueryactivityAPIRequest {
	return &AlibabawdkmarketingitembuygiftqueryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitembuygiftqueryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitembuygiftqueryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitembuygiftqueryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 查询入参
func (r *AlibabawdkmarketingitembuygiftqueryactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitembuygiftqueryactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
