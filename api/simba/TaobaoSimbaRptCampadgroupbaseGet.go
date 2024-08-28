package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCampadgroupbaseGet 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型)
// taobao.simba.rpt.campadgroupbase.get
//
// 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型)
func TaobaoSimbaRptCampadgroupbaseGet(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoSimbaRptCampadgroupbaseGetAPIRequest, resp *simba.TaobaoSimbaRptCampadgroupbaseGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
