package deliveryvoucher

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/deliveryvoucher"
)

// TaobaoGameDeliveryvoucherWatch 监控预约数据
// taobao.game.deliveryvoucher.watch
//
// 监控预约数据
func TaobaoGameDeliveryvoucherWatch(clt *core.SDKClient, req *deliveryvoucher.TaobaoGameDeliveryvoucherWatchAPIRequest, resp *deliveryvoucher.TaobaoGameDeliveryvoucherWatchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
