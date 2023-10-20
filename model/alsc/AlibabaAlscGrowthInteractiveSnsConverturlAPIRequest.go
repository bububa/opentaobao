package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscgrowthinteractivesnsconverturlAPIRequest 防封接口 API请求
// alibaba.alsc.growth.interactive.sns.converturl
//
// 防封接口
type AlibabaalscgrowthinteractivesnsconverturlAPIRequest struct {
	model.Params
	// 入参
	_param *FcUrlRequest
}

// NewAlibabaalscgrowthinteractivesnsconverturlRequest 初始化AlibabaalscgrowthinteractivesnsconverturlAPIRequest对象
func NewAlibabaalscgrowthinteractivesnsconverturlRequest() *AlibabaalscgrowthinteractivesnsconverturlAPIRequest {
	return &AlibabaalscgrowthinteractivesnsconverturlAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscgrowthinteractivesnsconverturlAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.growth.interactive.sns.converturl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscgrowthinteractivesnsconverturlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscgrowthinteractivesnsconverturlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaalscgrowthinteractivesnsconverturlAPIRequest) SetParam(_param *FcUrlRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaalscgrowthinteractivesnsconverturlAPIRequest) GetParam() *FcUrlRequest {
	return r._param
}
