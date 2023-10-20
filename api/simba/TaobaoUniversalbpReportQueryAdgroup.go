package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryAdgroup 单元报表查询
// taobao.universalbp.report.query.adgroup
//
// 单元报表查询
func TaobaoUniversalbpReportQueryAdgroup(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryAdgroupAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryAdgroupAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryAdgroupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
