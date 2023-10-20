package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkumsoutboundsortingcanclearareaAPIRequest dps分货-是否能够清场 API请求
// wdk.ums.outbound.sorting.cancleararea
//
// dps分货-是否能够清场
type WdkumsoutboundsortingcanclearareaAPIRequest struct {
	model.Params
	// 入参
	_param0 *DpsCanClearAreaMtopRequest
}

// NewWdkumsoutboundsortingcanclearareaRequest 初始化WdkumsoutboundsortingcanclearareaAPIRequest对象
func NewWdkumsoutboundsortingcanclearareaRequest() *WdkumsoutboundsortingcanclearareaAPIRequest {
	return &WdkumsoutboundsortingcanclearareaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkumsoutboundsortingcanclearareaAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.cancleararea"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkumsoutboundsortingcanclearareaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkumsoutboundsortingcanclearareaAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkumsoutboundsortingcanclearareaAPIRequest) SetParam0(_param0 *DpsCanClearAreaMtopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkumsoutboundsortingcanclearareaAPIRequest) GetParam0() *DpsCanClearAreaMtopRequest {
	return r._param0
}
