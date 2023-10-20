package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryRealtime 实时报表查询
// taobao.universalbp.report.query.realtime
//
// 实时报表查询
func TaobaoUniversalbpReportQueryRealtime(clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryRealtimeAPIRequest, session string) (*simba.TaobaoUniversalbpReportQueryRealtimeAPIResponse, error) {
	var resp simba.TaobaoUniversalbpReportQueryRealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
