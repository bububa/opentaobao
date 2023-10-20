package tmallhk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// Tmallhkclearanceget 天猫国际-清关材料查询
// tmall.hk.clearance.get
//
// 提供订单收货人身份信息查询功能。
func Tmallhkclearanceget(clt *core.SDKClient, req *tmallhk.TmallhkclearancegetAPIRequest, session string) (*tmallhk.TmallhkclearancegetAPIResponse, error) {
	var resp tmallhk.TmallhkclearancegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
