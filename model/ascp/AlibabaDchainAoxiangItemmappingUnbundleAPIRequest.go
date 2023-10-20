package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangitemmappingunbundleAPIRequest 商货关联解绑 API请求
// alibaba.dchain.aoxiang.itemmapping.unbundle
//
// 商货关联解绑
type AlibabadchainaoxiangitemmappingunbundleAPIRequest struct {
	model.Params
	// 商货关联解绑入参
	_unbundleItemmappingRequest *UnbundleItemmappingRequest
}

// NewAlibabadchainaoxiangitemmappingunbundleRequest 初始化AlibabadchainaoxiangitemmappingunbundleAPIRequest对象
func NewAlibabadchainaoxiangitemmappingunbundleRequest() *AlibabadchainaoxiangitemmappingunbundleAPIRequest {
	return &AlibabadchainaoxiangitemmappingunbundleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangitemmappingunbundleAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.itemmapping.unbundle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangitemmappingunbundleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangitemmappingunbundleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnbundleItemmappingRequest is UnbundleItemmappingRequest Setter
// 商货关联解绑入参
func (r *AlibabadchainaoxiangitemmappingunbundleAPIRequest) SetUnbundleItemmappingRequest(_unbundleItemmappingRequest *UnbundleItemmappingRequest) error {
	r._unbundleItemmappingRequest = _unbundleItemmappingRequest
	r.Set("unbundle_itemmapping_request", _unbundleItemmappingRequest)
	return nil
}

// GetUnbundleItemmappingRequest UnbundleItemmappingRequest Getter
func (r AlibabadchainaoxiangitemmappingunbundleAPIRequest) GetUnbundleItemmappingRequest() *UnbundleItemmappingRequest {
	return r._unbundleItemmappingRequest
}
