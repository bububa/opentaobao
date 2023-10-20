package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemmappingUnbundleAPIRequest 商货关联解绑 API请求
// alibaba.dchain.aoxiang.itemmapping.unbundle
//
// 商货关联解绑
type AlibabaDchainAoxiangItemmappingUnbundleAPIRequest struct {
	model.Params
	// 商货关联解绑入参
	_unbundleItemmappingRequest *UnbundleItemmappingRequest
}

// NewAlibabaDchainAoxiangItemmappingUnbundleRequest 初始化AlibabaDchainAoxiangItemmappingUnbundleAPIRequest对象
func NewAlibabaDchainAoxiangItemmappingUnbundleRequest() *AlibabaDchainAoxiangItemmappingUnbundleAPIRequest {
	return &AlibabaDchainAoxiangItemmappingUnbundleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangItemmappingUnbundleAPIRequest) Reset() {
	r._unbundleItemmappingRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangItemmappingUnbundleAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.itemmapping.unbundle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangItemmappingUnbundleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangItemmappingUnbundleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUnbundleItemmappingRequest is UnbundleItemmappingRequest Setter
// 商货关联解绑入参
func (r *AlibabaDchainAoxiangItemmappingUnbundleAPIRequest) SetUnbundleItemmappingRequest(_unbundleItemmappingRequest *UnbundleItemmappingRequest) error {
	r._unbundleItemmappingRequest = _unbundleItemmappingRequest
	r.Set("unbundle_itemmapping_request", _unbundleItemmappingRequest)
	return nil
}

// GetUnbundleItemmappingRequest UnbundleItemmappingRequest Getter
func (r AlibabaDchainAoxiangItemmappingUnbundleAPIRequest) GetUnbundleItemmappingRequest() *UnbundleItemmappingRequest {
	return r._unbundleItemmappingRequest
}

var poolAlibabaDchainAoxiangItemmappingUnbundleAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangItemmappingUnbundleRequest()
	},
}

// GetAlibabaDchainAoxiangItemmappingUnbundleRequest 从 sync.Pool 获取 AlibabaDchainAoxiangItemmappingUnbundleAPIRequest
func GetAlibabaDchainAoxiangItemmappingUnbundleAPIRequest() *AlibabaDchainAoxiangItemmappingUnbundleAPIRequest {
	return poolAlibabaDchainAoxiangItemmappingUnbundleAPIRequest.Get().(*AlibabaDchainAoxiangItemmappingUnbundleAPIRequest)
}

// ReleaseAlibabaDchainAoxiangItemmappingUnbundleAPIRequest 将 AlibabaDchainAoxiangItemmappingUnbundleAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangItemmappingUnbundleAPIRequest(v *AlibabaDchainAoxiangItemmappingUnbundleAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangItemmappingUnbundleAPIRequest.Put(v)
}
