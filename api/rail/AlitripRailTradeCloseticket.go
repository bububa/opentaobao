package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// AlitripRailTradeCloseticket 出票失败关单接口
// alitrip.rail.trade.closeticket
//
// 出票成功回调接口
func AlitripRailTradeCloseticket(clt *core.SDKClient, req *rail.AlitripRailTradeCloseticketAPIRequest, session string) (*rail.AlitripRailTradeCloseticketAPIResponse, error) {
	var resp rail.AlitripRailTradeCloseticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
