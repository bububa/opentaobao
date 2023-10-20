package mc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// Tmallmctaskchargelaunch 云码充电宝投放链路
// tmall.mc.task.charge.launch
//
// 云码充电宝投放链路，用于判断用户是否有匹配的投放计划
func Tmallmctaskchargelaunch(clt *core.SDKClient, req *mc.TmallmctaskchargelaunchAPIRequest, session string) (*mc.TmallmctaskchargelaunchAPIResponse, error) {
	var resp mc.TmallmctaskchargelaunchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
