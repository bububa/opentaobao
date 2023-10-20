package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoServindustryFinanceGeexOrderUpdate 即科订单结果更新回调
// taobao.servindustry.finance.geex.order.update
//
// 即科订单结果更新回调本地接口
func TaobaoServindustryFinanceGeexOrderUpdate(clt *core.SDKClient, req *trade.TaobaoServindustryFinanceGeexOrderUpdateAPIRequest, session string) (*trade.TaobaoServindustryFinanceGeexOrderUpdateAPIResponse, error) {
	var resp trade.TaobaoServindustryFinanceGeexOrderUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
