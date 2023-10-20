package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// Tmallmcdevicecirclecheck 云码设备圈选情况查询
// tmall.mc.device.circle.check
//
// 云码设备圈选情况查询
func Tmallmcdevicecirclecheck(clt *core.SDKClient, req *mc.TmallmcdevicecirclecheckAPIRequest, session string) (*mc.TmallmcdevicecirclecheckAPIResponse, error) {
	var resp mc.TmallmcdevicecirclecheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
