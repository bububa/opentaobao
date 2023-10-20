package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkumsoutboundsortinguserqueryAPIRequest dps-查询分货作业人员信息 API请求
// wdk.ums.outbound.sorting.userquery
//
// dps-查询分货作业人员信息
type WdkumsoutboundsortinguserqueryAPIRequest struct {
	model.Params
	// 入参
	_param0 *UserQueryByIdTopRequest
}

// NewWdkumsoutboundsortinguserqueryRequest 初始化WdkumsoutboundsortinguserqueryAPIRequest对象
func NewWdkumsoutboundsortinguserqueryRequest() *WdkumsoutboundsortinguserqueryAPIRequest {
	return &WdkumsoutboundsortinguserqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkumsoutboundsortinguserqueryAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.userquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkumsoutboundsortinguserqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkumsoutboundsortinguserqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkumsoutboundsortinguserqueryAPIRequest) SetParam0(_param0 *UserQueryByIdTopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkumsoutboundsortinguserqueryAPIRequest) GetParam0() *UserQueryByIdTopRequest {
	return r._param0
}
