package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkumsoutboundsortingcallbackforpulltaskAPIRequest dps分货-任务拉取确定接口 API请求
// wdk.ums.outbound.sorting.callbackforpulltask
//
// dps分货-任务拉取确定接口
type WdkumsoutboundsortingcallbackforpulltaskAPIRequest struct {
	model.Params
	// 入参
	_param0 *DpsCallBackForPullTaskMtopRequest
}

// NewWdkumsoutboundsortingcallbackforpulltaskRequest 初始化WdkumsoutboundsortingcallbackforpulltaskAPIRequest对象
func NewWdkumsoutboundsortingcallbackforpulltaskRequest() *WdkumsoutboundsortingcallbackforpulltaskAPIRequest {
	return &WdkumsoutboundsortingcallbackforpulltaskAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkumsoutboundsortingcallbackforpulltaskAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.callbackforpulltask"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkumsoutboundsortingcallbackforpulltaskAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkumsoutboundsortingcallbackforpulltaskAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkumsoutboundsortingcallbackforpulltaskAPIRequest) SetParam0(_param0 *DpsCallBackForPullTaskMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkumsoutboundsortingcallbackforpulltaskAPIRequest) GetParam0() *DpsCallBackForPullTaskMtopRequest {
	return r._param0
}
