package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderCreate 服务单创建
// tmall.servicecenter.spserviceorder.create
//
// 服务单创建
func TmallServicecenterSpserviceorderCreate(clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderCreateAPIRequest, session string) (*tmallservice.TmallServicecenterSpserviceorderCreateAPIResponse, error) {
	var resp tmallservice.TmallServicecenterSpserviceorderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
