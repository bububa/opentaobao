package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterReservecondUpdate 主动预约条件更新
// tmall.servicecenter.reservecond.update
//
// 1、设置主动预约开通条件
func TmallServicecenterReservecondUpdate(clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondUpdateAPIRequest, resp *tmallsc.TmallServicecenterReservecondUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
