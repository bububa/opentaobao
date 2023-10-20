package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallmsfreservation 喵师傅服务预约API
// tmall.msf.reservation
//
// 喵师傅预约api
func Tmallmsfreservation(clt *core.SDKClient, req *tmallservice.TmallmsfreservationAPIRequest, session string) (*tmallservice.TmallmsfreservationAPIResponse, error) {
	var resp tmallservice.TmallmsfreservationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
