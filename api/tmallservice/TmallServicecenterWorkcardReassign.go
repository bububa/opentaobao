package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardreassign 工单改派门店
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
func Tmallservicecenterworkcardreassign(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardreassignAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardreassignAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardreassignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
