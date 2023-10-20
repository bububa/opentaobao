package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkumssortingfullcontainerAPIRequest dps分货-满箱 API请求
// wdk.ums.sorting.full.container
//
// dps分货-满箱
type WdkumssortingfullcontainerAPIRequest struct {
	model.Params
	// 入参
	_param0 *DpsScanContainerMtopRequest
}

// NewWdkumssortingfullcontainerRequest 初始化WdkumssortingfullcontainerAPIRequest对象
func NewWdkumssortingfullcontainerRequest() *WdkumssortingfullcontainerAPIRequest {
	return &WdkumssortingfullcontainerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkumssortingfullcontainerAPIRequest) GetApiMethodName() string {
	return "wdk.ums.sorting.full.container"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkumssortingfullcontainerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkumssortingfullcontainerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkumssortingfullcontainerAPIRequest) SetParam0(_param0 *DpsScanContainerMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkumssortingfullcontainerAPIRequest) GetParam0() *DpsScanContainerMtopRequest {
	return r._param0
}
