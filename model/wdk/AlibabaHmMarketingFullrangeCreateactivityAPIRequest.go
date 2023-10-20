package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingfullrangecreateactivityAPIRequest 创建全场活动 API请求
// alibaba.hm.marketing.fullrange.createactivity
//
// 创建全场活动
type AlibabahmmarketingfullrangecreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *FullRangeActivity
}

// NewAlibabahmmarketingfullrangecreateactivityRequest 初始化AlibabahmmarketingfullrangecreateactivityAPIRequest对象
func NewAlibabahmmarketingfullrangecreateactivityRequest() *AlibabahmmarketingfullrangecreateactivityAPIRequest {
	return &AlibabahmmarketingfullrangecreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingfullrangecreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingfullrangecreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingfullrangecreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabahmmarketingfullrangecreateactivityAPIRequest) SetParam(_param *FullRangeActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingfullrangecreateactivityAPIRequest) GetParam() *FullRangeActivity {
	return r._param
}
