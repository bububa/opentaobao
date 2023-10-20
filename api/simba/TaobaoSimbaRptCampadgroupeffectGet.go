package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptcampadgroupeffectget 推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型)
// taobao.simba.rpt.campadgroupeffect.get
//
// 推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型)
func Taobaosimbarptcampadgroupeffectget(clt *core.SDKClient, req *simba.TaobaosimbarptcampadgroupeffectgetAPIRequest, session string) (*simba.TaobaosimbarptcampadgroupeffectgetAPIResponse, error) {
	var resp simba.TaobaosimbarptcampadgroupeffectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
