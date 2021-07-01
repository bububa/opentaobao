package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

/* TmallTraderateFeedsGet
查询子订单对应的评价、追评以及语义标签
tmall.traderate.feeds.get

通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签 */
func TmallTraderateFeedsGet(clt *core.SDKClient, req *traderate.TmallTraderateFeedsGetAPIRequest, session string) (*traderate.TmallTraderateFeedsGetAPIResponse, error) {
	var resp traderate.TmallTraderateFeedsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
