package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleOnespuRegisterUpdate 闲鱼 ONESPU 挂载接口
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
func AlibabaIdleOnespuRegisterUpdate(clt *core.SDKClient, req *idle.AlibabaIdleOnespuRegisterUpdateAPIRequest, session string) (*idle.AlibabaIdleOnespuRegisterUpdateAPIResponse, error) {
	var resp idle.AlibabaIdleOnespuRegisterUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
