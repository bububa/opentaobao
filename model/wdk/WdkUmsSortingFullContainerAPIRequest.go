package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsSortingFullContainerAPIRequest dps分货-满箱 API请求
// wdk.ums.sorting.full.container
//
// dps分货-满箱
type WdkUmsSortingFullContainerAPIRequest struct {
	model.Params
	// 入参
	_param0 *DpsScanContainerMtopRequest
}

// NewWdkUmsSortingFullContainerRequest 初始化WdkUmsSortingFullContainerAPIRequest对象
func NewWdkUmsSortingFullContainerRequest() *WdkUmsSortingFullContainerAPIRequest {
	return &WdkUmsSortingFullContainerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkUmsSortingFullContainerAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsSortingFullContainerAPIRequest) GetApiMethodName() string {
	return "wdk.ums.sorting.full.container"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsSortingFullContainerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkUmsSortingFullContainerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkUmsSortingFullContainerAPIRequest) SetParam0(_param0 *DpsScanContainerMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkUmsSortingFullContainerAPIRequest) GetParam0() *DpsScanContainerMtopRequest {
	return r._param0
}

var poolWdkUmsSortingFullContainerAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkUmsSortingFullContainerRequest()
	},
}

// GetWdkUmsSortingFullContainerRequest 从 sync.Pool 获取 WdkUmsSortingFullContainerAPIRequest
func GetWdkUmsSortingFullContainerAPIRequest() *WdkUmsSortingFullContainerAPIRequest {
	return poolWdkUmsSortingFullContainerAPIRequest.Get().(*WdkUmsSortingFullContainerAPIRequest)
}

// ReleaseWdkUmsSortingFullContainerAPIRequest 将 WdkUmsSortingFullContainerAPIRequest 放入 sync.Pool
func ReleaseWdkUmsSortingFullContainerAPIRequest(v *WdkUmsSortingFullContainerAPIRequest) {
	v.Reset()
	poolWdkUmsSortingFullContainerAPIRequest.Put(v)
}
