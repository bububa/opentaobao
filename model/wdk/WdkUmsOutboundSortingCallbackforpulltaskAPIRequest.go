package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingCallbackforpulltaskAPIRequest dps分货-任务拉取确定接口 API请求
// wdk.ums.outbound.sorting.callbackforpulltask
//
// dps分货-任务拉取确定接口
type WdkUmsOutboundSortingCallbackforpulltaskAPIRequest struct {
	model.Params
	// 入参
	_param0 *DpsCallBackForPullTaskMtopRequest
}

// NewWdkUmsOutboundSortingCallbackforpulltaskRequest 初始化WdkUmsOutboundSortingCallbackforpulltaskAPIRequest对象
func NewWdkUmsOutboundSortingCallbackforpulltaskRequest() *WdkUmsOutboundSortingCallbackforpulltaskAPIRequest {
	return &WdkUmsOutboundSortingCallbackforpulltaskAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *WdkUmsOutboundSortingCallbackforpulltaskAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsOutboundSortingCallbackforpulltaskAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.callbackforpulltask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsOutboundSortingCallbackforpulltaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkUmsOutboundSortingCallbackforpulltaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkUmsOutboundSortingCallbackforpulltaskAPIRequest) SetParam0(_param0 *DpsCallBackForPullTaskMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkUmsOutboundSortingCallbackforpulltaskAPIRequest) GetParam0() *DpsCallBackForPullTaskMtopRequest {
	return r._param0
}

var poolWdkUmsOutboundSortingCallbackforpulltaskAPIRequest = sync.Pool{
	New: func() any {
		return NewWdkUmsOutboundSortingCallbackforpulltaskRequest()
	},
}

// GetWdkUmsOutboundSortingCallbackforpulltaskRequest 从 sync.Pool 获取 WdkUmsOutboundSortingCallbackforpulltaskAPIRequest
func GetWdkUmsOutboundSortingCallbackforpulltaskAPIRequest() *WdkUmsOutboundSortingCallbackforpulltaskAPIRequest {
	return poolWdkUmsOutboundSortingCallbackforpulltaskAPIRequest.Get().(*WdkUmsOutboundSortingCallbackforpulltaskAPIRequest)
}

// ReleaseWdkUmsOutboundSortingCallbackforpulltaskAPIRequest 将 WdkUmsOutboundSortingCallbackforpulltaskAPIRequest 放入 sync.Pool
func ReleaseWdkUmsOutboundSortingCallbackforpulltaskAPIRequest(v *WdkUmsOutboundSortingCallbackforpulltaskAPIRequest) {
	v.Reset()
	poolWdkUmsOutboundSortingCallbackforpulltaskAPIRequest.Put(v)
}
