package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterReservecondCreate 创建主动预约开通条件
// tmall.servicecenter.reservecond.create
//
// 1、设置主动预约开通条件
func TmallServicecenterReservecondCreate(clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondCreateAPIRequest, resp *tmallsc.TmallServicecenterReservecondCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
