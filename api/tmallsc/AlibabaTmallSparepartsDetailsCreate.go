package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabatmallsparepartsdetailscreate 天猫蚁巢同步工单申请备件明细
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
func Alibabatmallsparepartsdetailscreate(clt *core.SDKClient, req *tmallsc.AlibabatmallsparepartsdetailscreateAPIRequest, session string) (*tmallsc.AlibabatmallsparepartsdetailscreateAPIResponse, error) {
	var resp tmallsc.AlibabatmallsparepartsdetailscreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
