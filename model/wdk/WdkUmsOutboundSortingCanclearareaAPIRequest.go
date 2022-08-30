package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingCanclearareaAPIRequest dps分货-是否能够清场 API请求
// wdk.ums.outbound.sorting.cancleararea
//
// dps分货-是否能够清场
type WdkUmsOutboundSortingCanclearareaAPIRequest struct {
	model.Params
	// 入参
	_param0 *DpsCanClearAreaMtopRequest
}

// NewWdkUmsOutboundSortingCanclearareaRequest 初始化WdkUmsOutboundSortingCanclearareaAPIRequest对象
func NewWdkUmsOutboundSortingCanclearareaRequest() *WdkUmsOutboundSortingCanclearareaAPIRequest {
	return &WdkUmsOutboundSortingCanclearareaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsOutboundSortingCanclearareaAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.cancleararea"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsOutboundSortingCanclearareaAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkUmsOutboundSortingCanclearareaAPIRequest) SetParam0(_param0 *DpsCanClearAreaMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkUmsOutboundSortingCanclearareaAPIRequest) GetParam0() *DpsCanClearAreaMtopRequest {
	return r._param0
}
