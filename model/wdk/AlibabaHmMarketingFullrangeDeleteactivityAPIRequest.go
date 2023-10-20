package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingfullrangedeleteactivityAPIRequest 全场活动删除活动接口 API请求
// alibaba.hm.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
type AlibabahmmarketingfullrangedeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabahmmarketingfullrangedeleteactivityRequest 初始化AlibabahmmarketingfullrangedeleteactivityAPIRequest对象
func NewAlibabahmmarketingfullrangedeleteactivityRequest() *AlibabahmmarketingfullrangedeleteactivityAPIRequest {
	return &AlibabahmmarketingfullrangedeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingfullrangedeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingfullrangedeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingfullrangedeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabahmmarketingfullrangedeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingfullrangedeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
