package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurcmallpackagesyncAPIRequest 套餐同步 API请求
// alibaba.pur.cmall.package.sync
//
// 套餐同步
type AlibabapurcmallpackagesyncAPIRequest struct {
	model.Params
	// 套餐对象
	_accessPackageDto *AccessPackageDto
}

// NewAlibabapurcmallpackagesyncRequest 初始化AlibabapurcmallpackagesyncAPIRequest对象
func NewAlibabapurcmallpackagesyncRequest() *AlibabapurcmallpackagesyncAPIRequest {
	return &AlibabapurcmallpackagesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapurcmallpackagesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.cmall.package.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapurcmallpackagesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapurcmallpackagesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessPackageDto is AccessPackageDto Setter
// 套餐对象
func (r *AlibabapurcmallpackagesyncAPIRequest) SetAccessPackageDto(_accessPackageDto *AccessPackageDto) error {
	r._accessPackageDto = _accessPackageDto
	r.Set("access_package_dto", _accessPackageDto)
	return nil
}

// GetAccessPackageDto AccessPackageDto Getter
func (r AlibabapurcmallpackagesyncAPIRequest) GetAccessPackageDto() *AccessPackageDto {
	return r._accessPackageDto
}
