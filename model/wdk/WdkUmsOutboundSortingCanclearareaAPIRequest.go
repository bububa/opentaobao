package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkUmsOutboundSortingCanclearareaAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsOutboundSortingCanclearareaAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.cancleararea"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsOutboundSortingCanclearareaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkUmsOutboundSortingCanclearareaAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolWdkUmsOutboundSortingCanclearareaAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkUmsOutboundSortingCanclearareaRequest()
	},
}

// GetWdkUmsOutboundSortingCanclearareaRequest 从 sync.Pool 获取 WdkUmsOutboundSortingCanclearareaAPIRequest
func GetWdkUmsOutboundSortingCanclearareaAPIRequest() *WdkUmsOutboundSortingCanclearareaAPIRequest {
	return poolWdkUmsOutboundSortingCanclearareaAPIRequest.Get().(*WdkUmsOutboundSortingCanclearareaAPIRequest)
}

// ReleaseWdkUmsOutboundSortingCanclearareaAPIRequest 将 WdkUmsOutboundSortingCanclearareaAPIRequest 放入 sync.Pool
func ReleaseWdkUmsOutboundSortingCanclearareaAPIRequest(v *WdkUmsOutboundSortingCanclearareaAPIRequest) {
	v.Reset()
	poolWdkUmsOutboundSortingCanclearareaAPIRequest.Put(v)
}
