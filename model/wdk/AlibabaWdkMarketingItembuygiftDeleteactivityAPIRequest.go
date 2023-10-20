package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest 删除买赠活动 API请求
// alibaba.wdk.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
type AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest struct {
	model.Params
	// 要删除的活动信息
	_param *CommonActivityParam
}

// NewAlibabawdkmarketingitembuygiftdeleteactivityRequest 初始化AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest对象
func NewAlibabawdkmarketingitembuygiftdeleteactivityRequest() *AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest {
	return &AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 要删除的活动信息
func (r *AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitembuygiftdeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
