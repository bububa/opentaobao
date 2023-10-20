package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptadgroupcreativebaseget 推广组下创意报表基础数据查询(汇总数据，不分类型)
// taobao.simba.rpt.adgroupcreativebase.get
//
// 推广组下创意报表基础数据查询(汇总数据，不分类型)
func Taobaosimbarptadgroupcreativebaseget(clt *core.SDKClient, req *simba.TaobaosimbarptadgroupcreativebasegetAPIRequest, session string) (*simba.TaobaosimbarptadgroupcreativebasegetAPIResponse, error) {
	var resp simba.TaobaosimbarptadgroupcreativebasegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
