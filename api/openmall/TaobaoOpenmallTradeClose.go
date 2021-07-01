package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

/* TaobaoOpenmallTradeClose
关闭订单
taobao.openmall.trade.close

关闭订单 */
func TaobaoOpenmallTradeClose(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeCloseAPIRequest, session string) (*openmall.TaobaoOpenmallTradeCloseAPIResponse, error) {
	var resp openmall.TaobaoOpenmallTradeCloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
