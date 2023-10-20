package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherOrdervoucher 预约接口
// taobao.game.deliveryvoucher.ordervoucher
//
// 提货券发券接口：同步券和订单的关联信息
func TaobaoGameDeliveryvoucherOrdervoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherOrdervoucherAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherOrdervoucherAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
