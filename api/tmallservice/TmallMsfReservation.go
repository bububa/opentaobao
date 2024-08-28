package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMsfReservation 喵师傅服务预约API
// tmall.msf.reservation
//
// 喵师傅预约api
func TmallMsfReservation(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallMsfReservationAPIRequest, resp *tmallservice.TmallMsfReservationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
