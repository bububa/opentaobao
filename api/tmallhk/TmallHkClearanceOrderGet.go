package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmallhkclearanceorderget 天猫国际订单清关信息
// tmall.hk.clearance.order.get
//
// 天猫国际订单清关信息
func Tmallhkclearanceorderget(clt *core.SDKClient, req *tmallhk.TmallhkclearanceordergetAPIRequest, session string) (*tmallhk.TmallhkclearanceordergetAPIResponse, error) {
	var resp tmallhk.TmallhkclearanceordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
