package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryCrowd 人群报表查询
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
func TaobaoUniversalbpReportQueryCrowd(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryCrowdAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryCrowdAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryCrowdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
