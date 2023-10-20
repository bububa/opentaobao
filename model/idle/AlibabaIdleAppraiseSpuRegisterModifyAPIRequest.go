package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleappraisespuregistermodifyAPIRequest 验货宝服务商spu挂载 API请求
// alibaba.idle.appraise.spu.register.modify
//
// 闲鱼接收回收商spu模板挂载信息
type AlibabaidleappraisespuregistermodifyAPIRequest struct {
	model.Params
	// 入参
	_idleAppraiseSpuRegister4TopDto *IdleAppraiseSpuRegister4topDto
}

// NewAlibabaidleappraisespuregistermodifyRequest 初始化AlibabaidleappraisespuregistermodifyAPIRequest对象
func NewAlibabaidleappraisespuregistermodifyRequest() *AlibabaidleappraisespuregistermodifyAPIRequest {
	return &AlibabaidleappraisespuregistermodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleappraisespuregistermodifyAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.spu.register.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleappraisespuregistermodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleappraisespuregistermodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdleAppraiseSpuRegister4TopDto is IdleAppraiseSpuRegister4TopDto Setter
// 入参
func (r *AlibabaidleappraisespuregistermodifyAPIRequest) SetIdleAppraiseSpuRegister4TopDto(_idleAppraiseSpuRegister4TopDto *IdleAppraiseSpuRegister4topDto) error {
	r._idleAppraiseSpuRegister4TopDto = _idleAppraiseSpuRegister4TopDto
	r.Set("idle_appraise_spu_register4_top_dto", _idleAppraiseSpuRegister4TopDto)
	return nil
}

// GetIdleAppraiseSpuRegister4TopDto IdleAppraiseSpuRegister4TopDto Getter
func (r AlibabaidleappraisespuregistermodifyAPIRequest) GetIdleAppraiseSpuRegister4TopDto() *IdleAppraiseSpuRegister4topDto {
	return r._idleAppraiseSpuRegister4TopDto
}
