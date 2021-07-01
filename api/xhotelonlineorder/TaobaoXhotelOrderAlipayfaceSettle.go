package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* TaobaoXhotelOrderAlipayfaceSettle
信用住订单结账接口
taobao.xhotel.order.alipayface.settle

用于离店付订单在客人离店后，发起结账以及扣款等后续动作 */
func TaobaoXhotelOrderAlipayfaceSettle(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderAlipayfaceSettleAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderAlipayfaceSettleAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelOrderAlipayfaceSettleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
