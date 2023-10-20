package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplancountryidgetAPIRequest 定向推广-国家标签ID获取 API请求
// alibaba.scbp.target.ad.plan.country.id.get
//
// 定向推广-国家标签ID获取
type AlibabascbptargetadplancountryidgetAPIRequest struct {
	model.Params
}

// NewAlibabascbptargetadplancountryidgetRequest 初始化AlibabascbptargetadplancountryidgetAPIRequest对象
func NewAlibabascbptargetadplancountryidgetRequest() *AlibabascbptargetadplancountryidgetAPIRequest {
	return &AlibabascbptargetadplancountryidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbptargetadplancountryidgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.target.ad.plan.country.id.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbptargetadplancountryidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbptargetadplancountryidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
