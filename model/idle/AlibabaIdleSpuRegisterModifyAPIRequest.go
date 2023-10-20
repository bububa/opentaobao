package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlespuregistermodifyAPIRequest 服务商spu挂载接口 API请求
// alibaba.idle.spu.register.modify
//
// 闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等)
type AlibabaidlespuregistermodifyAPIRequest struct {
	model.Params
	// 入参
	_idleSpuRegister4TopDto *IdleSpuRegister4topDto
}

// NewAlibabaidlespuregistermodifyRequest 初始化AlibabaidlespuregistermodifyAPIRequest对象
func NewAlibabaidlespuregistermodifyRequest() *AlibabaidlespuregistermodifyAPIRequest {
	return &AlibabaidlespuregistermodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlespuregistermodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.spu.register.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlespuregistermodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlespuregistermodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdleSpuRegister4TopDto is IdleSpuRegister4TopDto Setter
// 入参
func (r *AlibabaidlespuregistermodifyAPIRequest) SetIdleSpuRegister4TopDto(_idleSpuRegister4TopDto *IdleSpuRegister4topDto) error {
	r._idleSpuRegister4TopDto = _idleSpuRegister4TopDto
	r.Set("idle_spu_register4_top_dto", _idleSpuRegister4TopDto)
	return nil
}

// GetIdleSpuRegister4TopDto IdleSpuRegister4TopDto Getter
func (r AlibabaidlespuregistermodifyAPIRequest) GetIdleSpuRegister4TopDto() *IdleSpuRegister4topDto {
	return r._idleSpuRegister4TopDto
}
