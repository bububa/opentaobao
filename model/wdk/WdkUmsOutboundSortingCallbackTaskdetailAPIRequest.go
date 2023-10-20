package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingCallbackTaskdetailAPIRequest dps分货，明细回传 API请求
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
type WdkUmsOutboundSortingCallbackTaskdetailAPIRequest struct {
	model.Params
	// 参数
	_param0 *DpsCallBackSortMtopRequest
}

// NewWdkUmsOutboundSortingCallbackTaskdetailRequest 初始化WdkUmsOutboundSortingCallbackTaskdetailAPIRequest对象
func NewWdkUmsOutboundSortingCallbackTaskdetailRequest() *WdkUmsOutboundSortingCallbackTaskdetailAPIRequest {
	return &WdkUmsOutboundSortingCallbackTaskdetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.callback.taskdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) SetParam0(_param0 *DpsCallBackSortMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) GetParam0() *DpsCallBackSortMtopRequest {
	return r._param0
}

var poolWdkUmsOutboundSortingCallbackTaskdetailAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkUmsOutboundSortingCallbackTaskdetailRequest()
	},
}

// GetWdkUmsOutboundSortingCallbackTaskdetailRequest 从 sync.Pool 获取 WdkUmsOutboundSortingCallbackTaskdetailAPIRequest
func GetWdkUmsOutboundSortingCallbackTaskdetailAPIRequest() *WdkUmsOutboundSortingCallbackTaskdetailAPIRequest {
	return poolWdkUmsOutboundSortingCallbackTaskdetailAPIRequest.Get().(*WdkUmsOutboundSortingCallbackTaskdetailAPIRequest)
}

// ReleaseWdkUmsOutboundSortingCallbackTaskdetailAPIRequest 将 WdkUmsOutboundSortingCallbackTaskdetailAPIRequest 放入 sync.Pool
func ReleaseWdkUmsOutboundSortingCallbackTaskdetailAPIRequest(v *WdkUmsOutboundSortingCallbackTaskdetailAPIRequest) {
	v.Reset()
	poolWdkUmsOutboundSortingCallbackTaskdetailAPIRequest.Put(v)
}
