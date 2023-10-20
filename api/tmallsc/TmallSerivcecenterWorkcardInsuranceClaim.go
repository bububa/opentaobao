package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallserivcecenterworkcardinsuranceclaim 保险理赔回传工单记录
// tmall.serivcecenter.workcard.insurance.claim
//
// 保险理赔回传工单记录
func Tmallserivcecenterworkcardinsuranceclaim(clt *core.SDKClient, req *tmallsc.TmallserivcecenterworkcardinsuranceclaimAPIRequest, session string) (*tmallsc.TmallserivcecenterworkcardinsuranceclaimAPIResponse, error) {
	var resp tmallsc.TmallserivcecenterworkcardinsuranceclaimAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
