package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicestoreupdatestatus 网点/门店状态修改
// tmall.servicecenter.servicestore.updatestatus
//
// 修改网点/门店状态
func Tmallservicecenterservicestoreupdatestatus(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicestoreupdatestatusAPIRequest, session string) (*tmallservice.TmallservicecenterservicestoreupdatestatusAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicestoreupdatestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
