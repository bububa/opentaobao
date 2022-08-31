package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoTradeOrderskuUpdate 更新交易的销售属性
// taobao.trade.ordersku.update
//
// 只能更新发货前子订单的销售属性 &lt;br/&gt;只能更新价格相同的销售属性。对于拍下减库存的交易会同步更新销售属性的库存量。对于旺店的交易，要使用商品扩展信息中的SKU价格来比较。 &lt;br/&gt;必须使用sku_id或sku_props中的一个参数来更新，如果两个都传的话，sku_id优先
func TaobaoTradeOrderskuUpdate(clt *core.SDKClient, req *trade.TaobaoTradeOrderskuUpdateAPIRequest, session string) (*trade.TaobaoTradeOrderskuUpdateAPIResponse, error) {
	var resp trade.TaobaoTradeOrderskuUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
