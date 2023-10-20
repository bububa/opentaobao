package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthimriskqueryAPIRequest 问诊质控接口 API请求
// alibaba.alihealth.imrisk.query
//
// 阿里健康的问诊质控接口
type AlibabaalihealthimriskqueryAPIRequest struct {
	model.Params
	// 入参数
	_param0 *ImriskCheckCommand
}

// NewAlibabaalihealthimriskqueryRequest 初始化AlibabaalihealthimriskqueryAPIRequest对象
func NewAlibabaalihealthimriskqueryRequest() *AlibabaalihealthimriskqueryAPIRequest {
	return &AlibabaalihealthimriskqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthimriskqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.imrisk.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthimriskqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthimriskqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参数
func (r *AlibabaalihealthimriskqueryAPIRequest) SetParam0(_param0 *ImriskCheckCommand) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaalihealthimriskqueryAPIRequest) GetParam0() *ImriskCheckCommand {
	return r._param0
}
