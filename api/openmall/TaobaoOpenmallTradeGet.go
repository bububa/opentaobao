package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

/* TaobaoOpenmallTradeGet
查询订单详情
taobao.openmall.trade.get

查询订单详情 */
func TaobaoOpenmallTradeGet(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeGetAPIRequest, session string) (*openmall.TaobaoOpenmallTradeGetAPIResponse, error) {
	var resp openmall.TaobaoOpenmallTradeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
