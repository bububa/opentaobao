package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// Alitriptripvpagentordersearch 【国际机票】查询辅营订单列表
// alitrip.tripvp.agent.order.search
//
// 【国际机票】查询辅营订单列表
func Alitriptripvpagentordersearch(clt *core.SDKClient, req *ieagency.AlitriptripvpagentordersearchAPIRequest, session string) (*ieagency.AlitriptripvpagentordersearchAPIResponse, error) {
	var resp ieagency.AlitriptripvpagentordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
