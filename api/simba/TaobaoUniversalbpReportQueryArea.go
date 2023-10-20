package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryArea 地域报表查询
// taobao.universalbp.report.query.area
//
// 地域报表查询
func TaobaoUniversalbpReportQueryArea(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryAreaAPIRequest, resp *simba.TaobaoUniversalbpReportQueryAreaAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
