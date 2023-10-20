package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoServindustryFinanceGeexOrderUpdate 即科订单结果更新回调
// taobao.servindustry.finance.geex.order.update
//
// 即科订单结果更新回调本地接口
func TaobaoServindustryFinanceGeexOrderUpdate(clt *core.SDKClient, req *trade.TaobaoServindustryFinanceGeexOrderUpdateAPIRequest, resp *trade.TaobaoServindustryFinanceGeexOrderUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
