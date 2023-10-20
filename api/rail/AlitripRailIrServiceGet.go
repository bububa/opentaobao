package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// AlitripRailIrServiceGet 国际火车票仓位坐席查询
// alitrip.rail.ir.service.get
//
// 国际火车票标准仓位坐席查询
func AlitripRailIrServiceGet(clt *core.SDKClient, req *rail.AlitripRailIrServiceGetAPIRequest, session string) (*rail.AlitripRailIrServiceGetAPIResponse, error) {
	var resp rail.AlitripRailIrServiceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
