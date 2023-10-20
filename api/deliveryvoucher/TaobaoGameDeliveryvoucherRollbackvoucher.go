package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherRollbackvoucher 回滚券
// taobao.game.deliveryvoucher.rollbackvoucher
//
// 提货券发券接口：同步券和订单的关联信息
func TaobaoGameDeliveryvoucherRollbackvoucher(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherRollbackvoucherAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
