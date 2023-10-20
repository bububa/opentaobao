package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptargetadplancrowdidget 定向推广-人群标签ID获取(店铺老客、优选人群)
// alibaba.scbp.target.ad.plan.crowd.id.get
//
// 定向推广-人群标签ID获取(店铺老客、优选人群)
func Alibabascbptargetadplancrowdidget(clt *core.SDKClient, req *scbp.AlibabascbptargetadplancrowdidgetAPIRequest, session string) (*scbp.AlibabascbptargetadplancrowdidgetAPIResponse, error) {
	var resp scbp.AlibabascbptargetadplancrowdidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
