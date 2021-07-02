package cainiaoecc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEccExceptionsDelayCountAPIRequest 菜鸟控制塔包裹滞留异常统计信息获取 API请求
// cainiao.ecc.exceptions.delay.count
//
// 菜鸟控制塔包裹滞留异常统计信息获取
type CainiaoEccExceptionsDelayCountAPIRequest struct {
	model.Params
}

// NewCainiaoEccExceptionsDelayCountRequest 初始化CainiaoEccExceptionsDelayCountAPIRequest对象
func NewCainiaoEccExceptionsDelayCountRequest() *CainiaoEccExceptionsDelayCountAPIRequest {
	return &CainiaoEccExceptionsDelayCountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEccExceptionsDelayCountAPIRequest) GetApiMethodName() string {
	return "cainiao.ecc.exceptions.delay.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEccExceptionsDelayCountAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
