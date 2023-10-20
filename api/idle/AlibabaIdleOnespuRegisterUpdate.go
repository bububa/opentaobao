package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleonespuregisterupdate 闲鱼 ONESPU 挂载接口
// alibaba.idle.onespu.register.update
//
// 闲鱼 ONESPU 挂载接口
func Alibabaidleonespuregisterupdate(clt *core.SDKClient, req *idle.AlibabaidleonespuregisterupdateAPIRequest, session string) (*idle.AlibabaidleonespuregisterupdateAPIResponse, error) {
	var resp idle.AlibabaidleonespuregisterupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
