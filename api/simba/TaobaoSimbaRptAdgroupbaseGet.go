package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupbaseGet 推广组基础报表数据对象
// taobao.simba.rpt.adgroupbase.get
//
// 推广组基础报表数据对象
func TaobaoSimbaRptAdgroupbaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupbaseGetAPIRequest, resp *simba.TaobaoSimbaRptAdgroupbaseGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
