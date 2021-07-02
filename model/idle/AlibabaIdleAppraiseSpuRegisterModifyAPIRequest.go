package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseSpuRegisterModifyAPIRequest 验货宝服务商spu挂载 API请求
// alibaba.idle.appraise.spu.register.modify
//
// 闲鱼接收回收商spu模板挂载信息
type AlibabaIdleAppraiseSpuRegisterModifyAPIRequest struct {
	model.Params
	// 入参
	_idleAppraiseSpuRegister4TopDto *IdleAppraiseSpuRegister4TopDto
}

// NewAlibabaIdleAppraiseSpuRegisterModifyRequest 初始化AlibabaIdleAppraiseSpuRegisterModifyAPIRequest对象
func NewAlibabaIdleAppraiseSpuRegisterModifyRequest() *AlibabaIdleAppraiseSpuRegisterModifyAPIRequest {
	return &AlibabaIdleAppraiseSpuRegisterModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.spu.register.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIdleAppraiseSpuRegister4TopDto is IdleAppraiseSpuRegister4TopDto Setter
// 入参
func (r *AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) SetIdleAppraiseSpuRegister4TopDto(_idleAppraiseSpuRegister4TopDto *IdleAppraiseSpuRegister4TopDto) error {
	r._idleAppraiseSpuRegister4TopDto = _idleAppraiseSpuRegister4TopDto
	r.Set("idle_appraise_spu_register4_top_dto", _idleAppraiseSpuRegister4TopDto)
	return nil
}

// GetIdleAppraiseSpuRegister4TopDto IdleAppraiseSpuRegister4TopDto Getter
func (r AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) GetIdleAppraiseSpuRegister4TopDto() *IdleAppraiseSpuRegister4TopDto {
	return r._idleAppraiseSpuRegister4TopDto
}
