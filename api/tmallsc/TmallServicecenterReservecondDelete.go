package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterReservecondDelete 删除主动预约开通条件
// tmall.servicecenter.reservecond.delete
//
// 删除主动预约开通条件
func TmallServicecenterReservecondDelete(clt *core.SDKClient, req *tmallsc.TmallServicecenterReservecondDeleteAPIRequest, resp *tmallsc.TmallServicecenterReservecondDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
