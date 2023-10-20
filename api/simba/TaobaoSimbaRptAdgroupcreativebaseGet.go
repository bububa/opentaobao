package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptAdgroupcreativebaseGet 推广组下创意报表基础数据查询(汇总数据，不分类型)
// taobao.simba.rpt.adgroupcreativebase.get
//
// 推广组下创意报表基础数据查询(汇总数据，不分类型)
func TaobaoSimbaRptAdgroupcreativebaseGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptAdgroupcreativebaseGetAPIRequest, session string) (*simba.TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse, error) {
	var resp simba.TaobaoSimbaRptAdgroupcreativebaseGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
