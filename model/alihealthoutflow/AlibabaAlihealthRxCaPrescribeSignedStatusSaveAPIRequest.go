package alihealthoutflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest 处方ca认证 API请求
// alibaba.alihealth.rx.ca.prescribe.signed.status.save
//
// 处方ca认证
type AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest struct {
	model.Params
	// 入参实体
	_param0 *PrescribeStatusDto
}

// NewAlibabaalihealthrxcaprescribesignedstatussaveRequest 初始化AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest对象
func NewAlibabaalihealthrxcaprescribesignedstatussaveRequest() *AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest {
	return &AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.rx.ca.prescribe.signed.status.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参实体
func (r *AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest) SetParam0(_param0 *PrescribeStatusDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaalihealthrxcaprescribesignedstatussaveAPIRequest) GetParam0() *PrescribeStatusDto {
	return r._param0
}
