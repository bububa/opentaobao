package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// WdkUmsOutboundSortingUserqueryAPIRequest dps-查询分货作业人员信息 API请求
// wdk.ums.outbound.sorting.userquery
//
// dps-查询分货作业人员信息
type WdkUmsOutboundSortingUserqueryAPIRequest struct {
	model.Params
	// 入参
	_param0 *UserQueryByIdTopRequest
}

// NewWdkUmsOutboundSortingUserqueryRequest 初始化WdkUmsOutboundSortingUserqueryAPIRequest对象
func NewWdkUmsOutboundSortingUserqueryRequest() *WdkUmsOutboundSortingUserqueryAPIRequest {
	return &WdkUmsOutboundSortingUserqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r WdkUmsOutboundSortingUserqueryAPIRequest) GetApiMethodName() string {
	return "wdk.ums.outbound.sorting.userquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r WdkUmsOutboundSortingUserqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r WdkUmsOutboundSortingUserqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *WdkUmsOutboundSortingUserqueryAPIRequest) SetParam0(_param0 *UserQueryByIdTopRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r WdkUmsOutboundSortingUserqueryAPIRequest) GetParam0() *UserQueryByIdTopRequest {
	return r._param0
}
