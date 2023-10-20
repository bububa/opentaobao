package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportqueryrealtime 实时报表查询
// taobao.universalbp.report.query.realtime
//
// 实时报表查询
func Taobaouniversalbpreportqueryrealtime(clt *core.SDKClient, req *simba.TaobaouniversalbpreportqueryrealtimeAPIRequest, session string) (*simba.TaobaouniversalbpreportqueryrealtimeAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportqueryrealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
