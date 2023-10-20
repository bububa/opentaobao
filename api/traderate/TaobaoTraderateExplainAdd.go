package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

// TaobaoTraderateExplainAdd 商城评价解释接口
// taobao.traderate.explain.add
//
// 商城卖家给评价做出解释（买家追加评论后、评价超过30天的，都不能再做评价解释）
func TaobaoTraderateExplainAdd(clt *core.SDKClient, req *traderate.TaobaoTraderateExplainAddAPIRequest, resp *traderate.TaobaoTraderateExplainAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
