package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* TmallMsfReservation
喵师傅服务预约API
tmall.msf.reservation

喵师傅预约api */
func TmallMsfReservation(clt *core.SDKClient, req *tmallservice.TmallMsfReservationAPIRequest, session string) (*tmallservice.TmallMsfReservationAPIResponse, error) {
	var resp tmallservice.TmallMsfReservationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
