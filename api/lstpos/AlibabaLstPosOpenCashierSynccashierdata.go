package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// Alibabalstposopencashiersynccashierdata 收银快照同步接口(最多10条订单信息)
// alibaba.lst.pos.open.cashier.synccashierdata
//
// 收银快照同步接口(最多10条订单信息)
func Alibabalstposopencashiersynccashierdata(clt *core.SDKClient, req *lstpos.AlibabalstposopencashiersynccashierdataAPIRequest, session string) (*lstpos.AlibabalstposopencashiersynccashierdataAPIResponse, error) {
	var resp lstpos.AlibabalstposopencashiersynccashierdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
