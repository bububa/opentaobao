package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleOnespuRegisterUpdate 闲鱼 ONESPU 挂载接口
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
func AlibabaIdleOnespuRegisterUpdate(clt *core.SDKClient, req *idle.AlibabaIdleOnespuRegisterUpdateAPIRequest, resp *idle.AlibabaIdleOnespuRegisterUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
