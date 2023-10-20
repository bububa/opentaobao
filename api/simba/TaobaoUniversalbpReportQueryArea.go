package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryArea 地域报表查询
// taobao.universalbp.report.query.area
//
// 地域报表查询
func TaobaoUniversalbpReportQueryArea(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryAreaAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryAreaAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryAreaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
