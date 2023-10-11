package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeDeleteactivityAPIRequest 全场活动删除活动接口 API请求
// alibaba.hm.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
type AlibabaHmMarketingFullrangeDeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabaHmMarketingFullrangeDeleteactivityRequest 初始化AlibabaHmMarketingFullrangeDeleteactivityAPIRequest对象
func NewAlibabaHmMarketingFullrangeDeleteactivityRequest() *AlibabaHmMarketingFullrangeDeleteactivityAPIRequest {
	return &AlibabaHmMarketingFullrangeDeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingFullrangeDeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingFullrangeDeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingFullrangeDeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabaHmMarketingFullrangeDeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingFullrangeDeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
