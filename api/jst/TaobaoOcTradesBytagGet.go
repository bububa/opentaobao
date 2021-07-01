package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

/* TaobaoOcTradesBytagGet
标签查询订单
taobao.oc.trades.bytag.get

根据标签查询订单编号 */
func TaobaoOcTradesBytagGet(clt *core.SDKClient, req *jst.TaobaoOcTradesBytagGetAPIRequest, session string) (*jst.TaobaoOcTradesBytagGetAPIResponse, error) {
	var resp jst.TaobaoOcTradesBytagGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
