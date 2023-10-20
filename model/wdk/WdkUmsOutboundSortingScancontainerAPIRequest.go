package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingScancontainerAPIRequest dps分货-扫描分货容器判断是否可用 API请求
// wdk.ums.outbound.sorting.scancontainer
//
// dps分货-扫描分货容器判断是否可用
type WdkUmsOutboundSortingScancontainerAPIRequest struct {
	model.Params
	// 返回值
	_param0 *DpsScanContainerMtopRequest
}

// NewWdkUmsOutboundSortingScancontainerRequest 初始化WdkUmsOutboundSortingScancontainerAPIRequest对象
func NewWdkUmsOutboundSortingScancontainerRequest() *WdkUmsOutboundSortingScancontainerAPIRequest {
	return &WdkUmsOutboundSortingScancontainerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkUmsOutboundSortingScancontainerAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsOutboundSortingScancontainerAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.scancontainer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsOutboundSortingScancontainerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkUmsOutboundSortingScancontainerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 返回值
func (r *WdkUmsOutboundSortingScancontainerAPIRequest) SetParam0(_param0 *DpsScanContainerMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkUmsOutboundSortingScancontainerAPIRequest) GetParam0() *DpsScanContainerMtopRequest {
	return r._param0
}

var poolWdkUmsOutboundSortingScancontainerAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkUmsOutboundSortingScancontainerRequest()
	},
}

// GetWdkUmsOutboundSortingScancontainerRequest 从 sync.Pool 获取 WdkUmsOutboundSortingScancontainerAPIRequest
func GetWdkUmsOutboundSortingScancontainerAPIRequest() *WdkUmsOutboundSortingScancontainerAPIRequest {
	return poolWdkUmsOutboundSortingScancontainerAPIRequest.Get().(*WdkUmsOutboundSortingScancontainerAPIRequest)
}

// ReleaseWdkUmsOutboundSortingScancontainerAPIRequest 将 WdkUmsOutboundSortingScancontainerAPIRequest 放入 sync.Pool
func ReleaseWdkUmsOutboundSortingScancontainerAPIRequest(v *WdkUmsOutboundSortingScancontainerAPIRequest) {
	v.Reset()
	poolWdkUmsOutboundSortingScancontainerAPIRequest.Put(v)
}
