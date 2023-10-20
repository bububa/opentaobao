package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingfullrangequeryactivityAPIRequest 全场活动查询活动 API请求
// alibaba.hm.marketing.fullrange.queryactivity
//
// 全场活动查询活动
type AlibabahmmarketingfullrangequeryactivityAPIRequest struct {
	model.Params
	// 查询活动入参
	_param0 *CommonActivityParam
}

// NewAlibabahmmarketingfullrangequeryactivityRequest 初始化AlibabahmmarketingfullrangequeryactivityAPIRequest对象
func NewAlibabahmmarketingfullrangequeryactivityRequest() *AlibabahmmarketingfullrangequeryactivityAPIRequest {
	return &AlibabahmmarketingfullrangequeryactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingfullrangequeryactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.fullrange.queryactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingfullrangequeryactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingfullrangequeryactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询活动入参
func (r *AlibabahmmarketingfullrangequeryactivityAPIRequest) SetParam0(_param0 *CommonActivityParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingfullrangequeryactivityAPIRequest) GetParam0() *CommonActivityParam {
	return r._param0
}
