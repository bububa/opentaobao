package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleSpuRegisterModify 服务商spu挂载接口
// alibaba.idle.spu.register.modify
//
// 闲鱼服务商通过此接口进行spu挂载，指明自己支持对该spu的服务(回收、验货等)
func AlibabaIdleSpuRegisterModify(clt *core.SDKClient, req *idle.AlibabaIdleSpuRegisterModifyAPIRequest, resp *idle.AlibabaIdleSpuRegisterModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
