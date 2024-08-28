package mc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mc"
)

// TmallMcTaskChargeLaunch 云码充电宝投放链路
// tmall.mc.task.charge.launch
//
// 云码充电宝投放链路，用于判断用户是否有匹配的投放计划
func TmallMcTaskChargeLaunch(ctx context.Context, clt *core.SDKClient, req *mc.TmallMcTaskChargeLaunchAPIRequest, resp *mc.TmallMcTaskChargeLaunchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
