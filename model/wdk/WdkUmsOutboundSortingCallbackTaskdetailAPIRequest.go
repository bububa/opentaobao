package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkumsoutboundsortingcallbacktaskdetailAPIRequest dps分货，明细回传 API请求
// wdk.ums.outbound.sorting.callback.taskdetail
//
// dps分货-分货明细回传
type WdkumsoutboundsortingcallbacktaskdetailAPIRequest struct {
	model.Params
	// 参数
	_param0 *DpsCallBackSortMtopRequest
}

// NewWdkumsoutboundsortingcallbacktaskdetailRequest 初始化WdkumsoutboundsortingcallbacktaskdetailAPIRequest对象
func NewWdkumsoutboundsortingcallbacktaskdetailRequest() *WdkumsoutboundsortingcallbacktaskdetailAPIRequest {
	return &WdkumsoutboundsortingcallbacktaskdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkumsoutboundsortingcallbacktaskdetailAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.callback.taskdetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkumsoutboundsortingcallbacktaskdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkumsoutboundsortingcallbacktaskdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *WdkumsoutboundsortingcallbacktaskdetailAPIRequest) SetParam0(_param0 *DpsCallBackSortMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkumsoutboundsortingcallbacktaskdetailAPIRequest) GetParam0() *DpsCallBackSortMtopRequest {
	return r._param0
}
