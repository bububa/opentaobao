package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// CainiaoBmsOrderConsignConfirm BMS出库通知
// cainiao.bms.order.consign.confirm
//
// BMS出库后，通知ISV
func CainiaoBmsOrderConsignConfirm(clt *core.SDKClient, req *wlb.CainiaoBmsOrderConsignConfirmAPIRequest, resp *wlb.CainiaoBmsOrderConsignConfirmAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
