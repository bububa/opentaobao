package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

/* TaobaoJdsTradeTracesGet
获取单条订单跟踪详情
taobao.jds.trade.traces.get

获取聚石塔数据共享的交易全链路信息 */
func TaobaoJdsTradeTracesGet(clt *core.SDKClient, req *jst.TaobaoJdsTradeTracesGetAPIRequest, session string) (*jst.TaobaoJdsTradeTracesGetAPIResponse, error) {
	var resp jst.TaobaoJdsTradeTracesGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
