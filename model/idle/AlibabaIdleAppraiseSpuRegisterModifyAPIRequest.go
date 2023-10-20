package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) Reset() {
	r._idleAppraiseSpuRegister4TopDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.spu.register.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleAppraiseSpuRegisterModifyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleAppraiseSpuRegisterModifyRequest()
	},
}

// GetAlibabaIdleAppraiseSpuRegisterModifyRequest 从 sync.Pool 获取 AlibabaIdleAppraiseSpuRegisterModifyAPIRequest
func GetAlibabaIdleAppraiseSpuRegisterModifyAPIRequest() *AlibabaIdleAppraiseSpuRegisterModifyAPIRequest {
	return poolAlibabaIdleAppraiseSpuRegisterModifyAPIRequest.Get().(*AlibabaIdleAppraiseSpuRegisterModifyAPIRequest)
}

// ReleaseAlibabaIdleAppraiseSpuRegisterModifyAPIRequest 将 AlibabaIdleAppraiseSpuRegisterModifyAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleAppraiseSpuRegisterModifyAPIRequest(v *AlibabaIdleAppraiseSpuRegisterModifyAPIRequest) {
	v.Reset()
	poolAlibabaIdleAppraiseSpuRegisterModifyAPIRequest.Put(v)
}
