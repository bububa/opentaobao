package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// Cainiaobmsorderconsignconfirm BMS出库通知
// cainiao.bms.order.consign.confirm
//
// BMS出库后，通知ISV
func Cainiaobmsorderconsignconfirm(clt *core.SDKClient, req *wlb.CainiaobmsorderconsignconfirmAPIRequest, session string) (*wlb.CainiaobmsorderconsignconfirmAPIResponse, error) {
	var resp wlb.CainiaobmsorderconsignconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
