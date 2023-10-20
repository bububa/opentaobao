package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallserivcecenterservicerorderinsurancecallback 服务商回传保单信息
// tmall.serivcecenter.servicerorder.insurance.callback
//
// 服务商回传保单信息
func Tmallserivcecenterservicerorderinsurancecallback(clt *core.SDKClient, req *tmallsc.TmallserivcecenterservicerorderinsurancecallbackAPIRequest, session string) (*tmallsc.TmallserivcecenterservicerorderinsurancecallbackAPIResponse, error) {
	var resp tmallsc.TmallserivcecenterservicerorderinsurancecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
