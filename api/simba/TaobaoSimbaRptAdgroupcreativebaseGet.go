package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupcreativebaseGet 推广组下创意报表基础数据查询(汇总数据，不分类型)
// taobao.simba.rpt.adgroupcreativebase.get
//
// 推广组下创意报表基础数据查询(汇总数据，不分类型)
func TaobaoSimbaRptAdgroupcreativebaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupcreativebaseGetAPIRequest, resp *simba.TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
