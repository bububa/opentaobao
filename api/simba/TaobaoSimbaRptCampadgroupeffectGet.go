package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCampadgroupeffectGet 推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型)
// taobao.simba.rpt.campadgroupeffect.get
//
// 推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型)
func TaobaoSimbaRptCampadgroupeffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCampadgroupeffectGetAPIRequest, resp *simba.TaobaoSimbaRptCampadgroupeffectGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
