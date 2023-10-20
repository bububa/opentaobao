package cainiaoecc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoeccexceptionsdelaycountAPIRequest 菜鸟控制塔包裹滞留异常统计信息获取 API请求
// cainiao.ecc.exceptions.delay.count
//
// 菜鸟控制塔包裹滞留异常统计信息获取
type CainiaoeccexceptionsdelaycountAPIRequest struct {
	model.Params
}

// NewCainiaoeccexceptionsdelaycountRequest 初始化CainiaoeccexceptionsdelaycountAPIRequest对象
func NewCainiaoeccexceptionsdelaycountRequest() *CainiaoeccexceptionsdelaycountAPIRequest {
	return &CainiaoeccexceptionsdelaycountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoeccexceptionsdelaycountAPIRequest) GetApiMethodName() string {
	return "cainiao.ecc.exceptions.delay.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoeccexceptionsdelaycountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoeccexceptionsdelaycountAPIRequest) GetRawParams() model.Params {
	return r.Params
}
