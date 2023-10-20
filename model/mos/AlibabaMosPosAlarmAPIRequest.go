package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosposalarmAPIRequest pos故障报警 API请求
// alibaba.mos.pos.alarm
//
// 故障报警
type AlibabamosposalarmAPIRequest struct {
	model.Params
	// 参数
	_param0 *PosLogDto
}

// NewAlibabamosposalarmRequest 初始化AlibabamosposalarmAPIRequest对象
func NewAlibabamosposalarmRequest() *AlibabamosposalarmAPIRequest {
	return &AlibabamosposalarmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosposalarmAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.pos.alarm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosposalarmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosposalarmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabamosposalarmAPIRequest) SetParam0(_param0 *PosLogDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabamosposalarmAPIRequest) GetParam0() *PosLogDto {
	return r._param0
}
