package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

/* CainiaoBmsOrderConsignConfirm
BMS出库通知
cainiao.bms.order.consign.confirm

BMS出库后，通知ISV */
func CainiaoBmsOrderConsignConfirm(clt *core.SDKClient, req *wlb.CainiaoBmsOrderConsignConfirmAPIRequest, session string) (*wlb.CainiaoBmsOrderConsignConfirmAPIResponse, error) {
	var resp wlb.CainiaoBmsOrderConsignConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
