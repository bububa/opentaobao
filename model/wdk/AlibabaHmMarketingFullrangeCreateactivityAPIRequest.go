package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingFullrangeCreateactivityAPIRequest 创建全场活动 API请求
// alibaba.hm.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabaHmMarketingFullrangeCreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *FullRangeActivity
}

// NewAlibabaHmMarketingFullrangeCreateactivityRequest 初始化AlibabaHmMarketingFullrangeCreateactivityAPIRequest对象
func NewAlibabaHmMarketingFullrangeCreateactivityRequest() *AlibabaHmMarketingFullrangeCreateactivityAPIRequest {
	return &AlibabaHmMarketingFullrangeCreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingFullrangeCreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingFullrangeCreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingFullrangeCreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaHmMarketingFullrangeCreateactivityAPIRequest) SetParam(_param *FullRangeActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingFullrangeCreateactivityAPIRequest) GetParam() *FullRangeActivity {
	return r._param
}
