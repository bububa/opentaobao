package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleSpuRegisterModifyAPIRequest
服务商spu挂载接口 API请求
alibaba.idle.spu.register.modify

闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等) */
type AlibabaIdleSpuRegisterModifyAPIRequest struct {
	model.Params
	// 入参
	_idleSpuRegister4TopDto *IdleSpuRegister4TopDto
}

// New
