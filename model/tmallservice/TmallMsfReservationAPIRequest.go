package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMsfReservationAPIRequest
喵师傅服务预约API API请求
tmall.msf.reservation

喵师傅预约api */
type TmallMsfReservationAPIRequest struct {
	model.Params
	// 预约内容
	_reservInfo *ReservationDto
}

// New
