package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
