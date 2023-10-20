package pur

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCmallPackageSyncAPIRequest 套餐同步 API请求
// alibaba.pur.cmall.package.sync
//
// 套餐同步
type AlibabaPurCmallPackageSyncAPIRequest struct {
	model.Params
	// 套餐对象
	_accessPackageDto *AccessPackageDto
}

// NewAlibabaPurCmallPackageSyncRequest 初始化AlibabaPurCmallPackageSyncAPIRequest对象
func NewAlibabaPurCmallPackageSyncRequest() *AlibabaPurCmallPackageSyncAPIRequest {
	return &AlibabaPurCmallPackageSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPurCmallPackageSyncAPIRequest) Reset() {
	r._accessPackageDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPurCmallPackageSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.cmall.package.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPurCmallPackageSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPurCmallPackageSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessPackageDto is AccessPackageDto Setter
// 套餐对象
func (r *AlibabaPurCmallPackageSyncAPIRequest) SetAccessPackageDto(_accessPackageDto *AccessPackageDto) error {
	r._accessPackageDto = _accessPackageDto
	r.Set("access_package_dto", _accessPackageDto)
	return nil
}

// GetAccessPackageDto AccessPackageDto Getter
func (r AlibabaPurCmallPackageSyncAPIRequest) GetAccessPackageDto() *AccessPackageDto {
	return r._accessPackageDto
}

var poolAlibabaPurCmallPackageSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPurCmallPackageSyncRequest()
	},
}

// GetAlibabaPurCmallPackageSyncRequest 从 sync.Pool 获取 AlibabaPurCmallPackageSyncAPIRequest
func GetAlibabaPurCmallPackageSyncAPIRequest() *AlibabaPurCmallPackageSyncAPIRequest {
	return poolAlibabaPurCmallPackageSyncAPIRequest.Get().(*AlibabaPurCmallPackageSyncAPIRequest)
}

// ReleaseAlibabaPurCmallPackageSyncAPIRequest 将 AlibabaPurCmallPackageSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaPurCmallPackageSyncAPIRequest(v *AlibabaPurCmallPackageSyncAPIRequest) {
	v.Reset()
	poolAlibabaPurCmallPackageSyncAPIRequest.Put(v)
}
