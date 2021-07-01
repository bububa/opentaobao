package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* TaobaoWttTradeServiceGet
获取网厅号卡垂直标信息
taobao.wtt.trade.service.get

查询网厅订单信息 */
func TaobaoWttTradeServiceGet(clt *core.SDKClient, req *trade.TaobaoWttTradeServiceGetAPIRequest, session string) (*trade.TaobaoWttTradeServiceGetAPIResponse, error) {
	var resp trade.TaobaoWttTradeServiceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
