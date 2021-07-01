package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAppraiseSpuRegisterModifyAPIRequest
验货宝服务商spu挂载 API请求
alibaba.idle.appraise.spu.register.modify

闲鱼接收回收商spu模板挂载信息 */
type AlibabaIdleAppraiseSpuRegisterModifyAPIRequest struct {
	model.Params
	// 入参
	_idleAppraiseSpuRegister4TopDto *IdleAppraiseSpuRegister4TopDto
}

// New
